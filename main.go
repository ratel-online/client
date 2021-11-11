package main

import (
	"github.com/ratel-online/client/shell"
	"os"
)

func main() {
	os.Args = append(os.Args, "ainililia@163.com")
	shell.New("127.0.0.1:8080").Start()
}
