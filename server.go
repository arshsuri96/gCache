package main

import "arshsuri96/ggcache/cache"

type ServerOpt struct {
	IsLeader   bool
	ListenAddr string
}

type Server struct {
	ServerOpt
	cache cache.Cacher
}
