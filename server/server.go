package server

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/element-of/harmond/caps"

	irc "gopkg.in/irc.v1"
)

type Server struct {
	Name string

	smutex   *sync.RWMutex
	sessions map[net.Conn]*Session

	cmutex  *sync.RWMutex
	clients map[string]*Client
	nicks   map[string]*Client
}

type Session struct {
	sock net.Conn

	isRegistered bool
	password     string
	clientID     string

	nick, user, real, host, cloak string

	capSet *caps.CapSet

	outq chan *irc.Message
}

// Client is a fully registered client on the network
type Client struct {
	sessions []*Session
}

func NewServer(name string) *Server {
	return &Server{
		Name: name,

		smutex:   &sync.RWMutex{},
		sessions: map[net.Conn]*Session{},

		cmutex:  &sync.RWMutex{},
		clients: map[string]*Client{},
		nicks:   map[string]*Client{},
	}
}

func (s *Server) Serve(l net.Listener) {
	for {
		cli, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go s.handleConn(cli)
	}
}

func (s *Server) handleConn(cli net.Conn) {
	se := &Session{
		sock:   cli,
		capSet: &caps.CapSet{},
		outq:   make(chan *irc.Message, 50),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		en := irc.NewWriter(cli)

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-se.outq:
				en.WriteMessage(msg)
			}
		}
	}()

	s.smutex.Lock()
	s.sessions[cli] = se
	s.smutex.Unlock()

	ircReader := irc.NewReader(cli)

	for {
		line, err := ircReader.ReadMessage()
		if err != nil {
			switch err {
			case irc.ErrZeroLengthMessage, irc.ErrMissingDataAfterPrefix, irc.ErrMissingDataAfterTags, irc.ErrMissingCommand:
				// Clients messing up the formatting of IRC lines will have irc lines dropped silently.
				continue
			default:
				log.Printf("client %v: %v", cli.RemoteAddr().String(), err)
				goto disconnect
			}
		}

		log.Printf("%#v", line)

		switch line.Command {
		case "NICK":
			if len(line.Params) != 1 {
				continue
			}

			nick := line.Params[0]
			se.nick = nick

		case "USER":
			if len(line.Params) != 4 {
				continue
			}

			ident := line.Params[0]
			real := line.Params[3]

			se.user = ident
			se.real = real

			// TODO: move this into (*Server).register(*Session)
			// TODO: check nick in use, etc there too
			se.isRegistered = true

			se.outq <- &irc.Message{
				Prefix: &irc.Prefix{
					Name: s.Name,
				},
				Command: "001",
				Params:  []string{se.nick, "Welcome to the Internet Relay Chat network"},
			}

		case "PING":
			line.Command = "PONG"
			line.Prefix.Name = s.Name
			se.outq <- line
		}
	}

disconnect:
	s.smutex.Lock()
	delete(s.sessions, cli)
	s.smutex.Unlock()
	cli.Close()
}
