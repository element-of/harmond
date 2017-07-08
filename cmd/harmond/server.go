package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Xe/ln"
	"github.com/element-of/harmond/internal"
	"github.com/element-of/harmond/internal/numeric"
	cmap "github.com/orcaman/concurrent-map"
	"gopkg.in/irc.v1"
)

// NickSort normalizes a nickname for sorting and map storage.
func NickSort(nick string) string {
	return strings.ToLower(nick)
}

// Client holds information about a client on the IRC network.
type Client struct {
	sync.Mutex
	sock       net.Conn
	registered bool
	metadata   map[string]string
	outq       chan *irc.Message
	cancel     context.CancelFunc

	Nick        string
	User        string
	Host        string
	VHost       string
	IP          string
	Account     string
	ID          string
	Gecos       string
	Permissions int
	Umodes      int
	TS          time.Time
	LastSeen    time.Time
	Channels    []string
	Certfp      string
}

func (c *Client) SendLine(m *irc.Message) {
	if l := len(c.outq); l >= 50 {
		ln.Log(c, ln.F{"action": "*Client.SendLine", "msg": "client has large outq", "len": l})
	}
	c.outq <- m
}

// F ields for logging
func (c *Client) F() ln.F {
	return ln.F{
		"client_nick":       c.Nick,
		"client_account":    c.Account,
		"client_ts":         c.TS,
		"client_ip":         c.IP,
		"client_id":         c.ID,
		"client_certfp":     c.Certfp,
		"client_registered": c.registered,
		"client_last_seen":  c.LastSeen,
	}
}

// Config is the server configuration.
type Config struct {
	ServerName  string `env:"SERVER_NAME,required"`
	ServerID    string `env:"SERVER_ID"`
	ServerGecos string `env:"SERVER_GECOS"`
	NetworkName string `env:"NETWORK_NAME,required"`

	SSLPorts          []string `env:"SSL_PORTS"`
	PlainPorts        []string `env:"PLAIN_PORTS"`
	WebsocketSSLPorts []string `env:"WEBSOCKET_SSL_PORTS"`
}

// Server is the parent IRC server structure
type Server struct {
	// the list of channels the server knows about
	channels cmap.ConcurrentMap
	// the list of clients the server knows about
	clients *ClientSet
	// nextuid is the next UID that the server will generate for users.
	nextuid int
	// server configuration
	cfg Config
}

func newServer(cfg Config) *Server {
	return &Server{
		channels: cmap.New(),
		clients:  newClientSet(),
		nextuid:  100000,
		cfg:      cfg,
	}
}

// NextUID returns the next unique ID for a newly connected client.
func (s *Server) NextUID() string {
	s.nextuid++
	return s.cfg.ServerID + strconv.Itoa(s.nextuid)
}

// SendNumeric sends an IRC numeric to a client.
//
// See RFC1459 Section 2.4 for more information.
func (s *Server) SendNumeric(c *Client, n numeric.Response, args ...interface{}) {
	c.SendLine(&irc.Message{
		Prefix: &irc.Prefix{
			Host: s.cfg.ServerName,
		},
		Command: string(n),
		Params:  n.FormStr(args...),
	})
}

func (s *Server) Kill(c *Client, why string) {
	c.SendLine(&irc.Message{
		Command: "ERROR",
		Params:  []string{c.Nick, why},
	})
	ln.Log(c, ln.F{"action": "server_kill", "reason": why})
	c.cancel()
	c.sock.Close()
}

func (s *Server) PreRegHandler(c *Client, line *irc.Message) bool {
	switch line.Command {
	case "NICK":
		if len(line.Params) != 1 {
			s.SendNumeric(c, numeric.ErrNonicknamegiven, "*")
		}

		nick := line.Params[0]

		if _, ok := s.clients.GetNick(nick); ok {
			s.SendNumeric(c, numeric.ErrNicknameinuse, "*", nick)
			return false
		}

		c.Nick = nick

	case "USER":
		if len(line.Params) != 4 {
			return false
		}

		// TODO: validate and set these
		ident := line.Params[0]
		real := line.Params[3]

		if !internal.IsAlpha(ident) {
			s.Kill(c, "Invalid ident, ident may only contain alphabetical characters")
		}

		c.User = ident
		c.Gecos = real

	case "PING":
		line.Command = "PONG"
		line.Prefix.Name = s.cfg.ServerName
		c.SendLine(line)
	}

	if c.Nick != "" && c.User != "" && c.Gecos != "" {
		return true
	}

	return false
}

// Serve makes the server infinitely listen on a given net.Listener.
func (s *Server) Serve(l net.Listener) {
	for {
		cli, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go s.Handle(cli, nil)
	}
}

// Handle blocks infinitely with a unique `net.Conn`.
func (s *Server) Handle(conn net.Conn, cs *tls.ConnectionState) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	h, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	if h == "" {
		h = "127.0.0.1"
	}

	c := &Client{
		metadata: map[string]string{},
		outq:     make(chan *irc.Message, 50),
		cancel:   cancel,
		sock:     conn,

		TS:      time.Now(),
		ID:      s.NextUID(),
		IP:      h,
		Account: "*",
	}
	defer s.clients.Del(c)

	go func() {
		en := irc.NewWriter(conn)

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-c.outq:
				err := en.WriteMessage(msg)
				if err != nil {
					ln.Error(err, c, ln.F{"action": "en.WriteMessage"})
					cancel()
				}
			}
		}
	}()

	ircReader := irc.NewReader(conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		line, err := ircReader.ReadMessage()
		if err != nil {
			switch err {
			case irc.ErrZeroLengthMessage, irc.ErrMissingDataAfterPrefix, irc.ErrMissingDataAfterTags, irc.ErrMissingCommand:
				// Clients messing up the formatting of IRC lines will have irc lines dropped silently.
				continue
			default:
				ln.Error(err, c, ln.F{"action": "ircReader.ReadMessage"})
				return
			}
		}

		c.LastSeen = time.Now()

		if !c.registered {
			if s.PreRegHandler(c, line) {
				c.registered = true
				ln.Log(c, ln.F{"action": "client_registered"})
				// Send welcome
			}
		}
	}
}
