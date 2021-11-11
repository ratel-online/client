package shell

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/ratel-online/client/api"
	"github.com/ratel-online/client/ctx"
	"log"
	"os"
)

var servers = []string{
	"https://raw.githubusercontent.com/awesome-cmd/chat/main/servers.json",
	"https://gitee.com/ainilili/chat/raw/main/servers.json",
}

type shell struct {
	ctx  *ctx.Context
	addr string
}

func New(addr string) *shell {
	return &shell{
		ctx:  &ctx.Context{},
		addr: addr,
	}
}

func (s *shell) Start() {
	args := os.Args
	if len(args) < 1 {
		log.Fatalln("please enter your username.")
	}
	username := args[1]
	fmt.Printf("password: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := api.Login(username, string(password))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Data.Token)
}
