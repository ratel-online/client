package main

import (
	"flag"
	"fmt"
	"github.com/ratel-online/client/shell"
	"github.com/ratel-online/core/log"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.IntVar(&port, "p", 9999, "port")
	flag.Parse()
}

func main() {
	log.Error(shell.New(fmt.Sprintf("%s:%d", host, port)).Start())
}
