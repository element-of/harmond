package main

import (
	"github.com/Xe/ln"
	"github.com/caarlos0/env"
)

func main() {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		ln.FatalErr(err, ln.F{"action": "env.Parse"})
	}
}
