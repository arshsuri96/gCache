package main

import (
	"arshsuri96/ggcache/cache"
	"context"
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
		go s.handleCommand(conn, buf[:n])
	}
}

func (s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
	msg, err := parseMessage(rawCmd)
	if err != nil {
		fmt.Println("failed to parse commands ", err)
		return
	}

	switch msg.Cmd {
	case CMDSet:
		err = s.handleSetCmd(conn, msg)
	case CMDGet:
		err = s.handleGetCmd(conn, msg)
	}
}

func (s *Server) handleGetCmd(conn net.Conn, msg *Message) error {
	val, err := s.cache.Get(msg.Key)
	if err != nil {
		return err
	}
	_, err = conn.Write(val)

	return err
}

func (s *Server) handleSetCmd(conn net.Conn, msg *Message) error {
	if err := s.cache.Set(msg.Key, msg.Value, msg.TTL); err != nil {
		return err
	}

	go s.sendToFollowers(context.TODO(), msg)

	return nil
}
