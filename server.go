package main

import (
	"arshsuri96/ggcache/cache"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
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
		go s.handleCommand(conn, buf[:n])
	}
}

func (s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
	rawStr := string(rawCmd)
	parts := strings.Split(rawStr, " ")
	if len(parts) == 0 {
		log.Println("invalid command")
		return
	}
	cmd := Command(parts[0])
	if cmd == CMDSet {
		if len(parts) != 4 {
			log.Println("invalid SET commands")
			return
		}
		ttl, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Println("invalid SET")
			return
		}

		msg := MSGSet{
			key:   []byte(parts[1]),
			value: []byte(parts[2]),
			TTL:   time.Duration(ttl),
		}
	}

}
