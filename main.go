package main

import (
	"github.com/ratel-online/client/shell"
	"github.com/ratel-online/core/log"
)

func main() {
	log.Error(shell.New("127.0.0.1:9999").Start())
}
