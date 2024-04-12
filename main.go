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
	name string
)

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "host")
	flag.IntVar(&port, "p", 9999, "port")
	flag.StringVar(&name, "n", "", "name")
	flag.Parse()
}

func main() {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Error(shell.New(addr, name).Start())
}
