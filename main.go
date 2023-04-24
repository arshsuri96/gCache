package main

import (
	"arshsuri96/ggcache/cache"
	"log"
	"net"
	"time"
)

func main() {
	opts := ServerOpts{
		ListenAddr: "3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("SET FOO BAR 2500"))
	}()

	server := NewServer(opts, cache.New())
	server.Start()

}
