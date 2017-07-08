package main

import (
	"net"

	"github.com/Xe/ln"
	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		ln.FatalErr(err, ln.F{"action": "env.Parse"})
	}

	s := newServer(cfg)

	for _, p := range cfg.PlainPorts {
		l, err := net.Listen("tcp", ":"+p)
		if err != nil {
			ln.FatalErr(err, ln.F{"action": "net.Listen", "port": p})
		}

		go s.Serve(l)
	}

	select {}
}
