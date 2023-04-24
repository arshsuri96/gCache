package main

import (
	"arshsuri96/ggcache/cache"
	"fmt"
	"log"
	"net"
)

type ServerOpts struct {
	IsLeader   bool
	ListenAddr string
}

type Server struct {
	ServerOpts
	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		fmt.Errorf("listening error %s", err)
	}

	log.Printf("listening on port [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("accept errors: %s\n", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read error %s", err)
			break
		}
		msg := buf[:n]
		fmt.Printf(string(msg))
	}
}
