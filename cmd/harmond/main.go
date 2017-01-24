package main

import (
	"flag"
	"log"
	"net"

	"github.com/element-of/harmond/server"
)

func main() {
	flag.Parse()

	s := server.NewServer("irc.harmo.nd")

	l, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}

	s.Serve(l)
}
