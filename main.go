package main

import (
	"github.com/ratel-online/client/shell"
)

func main() {
	shell.New("127.0.0.1:8080").Start()
}
