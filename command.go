package main

import "time"

type Command string

const (
	CMDSet Command = "SET"
	CMDGet Command = "GET"
)

type MSGSet struct {
	key   []byte
	value []byte
	TTL   time.Duration
}

type MSGGet struct {
	key []byte
}
