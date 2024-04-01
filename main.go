package main

import (
	"flag"
	"fmt"
	"github.com/ratel-online/client/shell"
	"github.com/ratel-online/core/log"
)

var serverHost string
var serverPort uint

func init() {
	flag.StringVar(&serverHost, "host", "127.0.0.1", "The landlords server host")
	flag.UintVar(&serverPort, "port", 9999, "The landlords server port")
}

func main() {
	flag.Parse()
	serverAddress := fmt.Sprintf("%s:%d", serverHost, serverPort)
	log.Error(shell.New(serverAddress).Start())
}
