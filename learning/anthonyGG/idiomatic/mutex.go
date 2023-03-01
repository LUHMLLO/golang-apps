package main

import (
	"net"
	"sync"
)

type Server struct {
	listenAddr string
	isRunning  bool

	mu    sync.RWMutex
	peers map[string]net.Conn

	otherMu sync.RWMutex
	other   map[string]net.Conn
}
