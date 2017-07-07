package main

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Xe/ln"
	cmap "github.com/orcaman/concurrent-map"
)

// NickSort normalizes a nickname for sorting and map storage.
func NickSort(nick string) string {
	return strings.ToLower(nick)
}

// Client holds information about a client on the IRC network.
type Client struct {
	*sync.Mutex
	sock net.Conn

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
	Channels    []string
	Certfp      string
	Metadata    map[string]string
}

// F ields for logging
func (c Client) F() ln.F {
	return ln.F{
		"client_nick":    c.Nick,
		"client_account": c.Account,
		"client_ts":      c.TS,
		"client_ip":      c.IP,
		"client_id":      c.ID,
		"client_certfp":  c.Certfp,
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
