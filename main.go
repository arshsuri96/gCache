package main

import (
	"arshsuri96/ggcache/cache"
	"flag"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("SET FOO Bar 40000"))
	if err != nil {
		log.Fatal(err)
	}

	var (
		listenAddr = flag.String("listenAddr", ":3000", "listen address of the server")
		leaderAddr = flag.String("leaderaddr", "", "listen address of the leader")
	)

	flag.Parse()

	opts := ServerOpts{
		ListenAddr: *listenAddr,
		IsLeader:   len(*leaderAddr) == 0,
		LeaderAddr: *leaderAddr,
	}

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	conn, err := net.Dial("tcp", ":3000")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	conn.Write([]byte("SET FOO BAR 2500000000"))
	// 	time.Sleep(time.Second * 2)

	// 	buf := make([]byte, 1000)
	// 	n, _ := conn.Read(buf)
	// 	fmt.Println((string(buf[:n])))

	// }()

	server := NewServer(opts, cache.New())
	server.Start()

}
