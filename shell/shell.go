package shell

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/ratel-online/client/api"
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/util"
	"github.com/ratel-online/core/log"
)

type shell struct {
	ctx  *ctx.Context
	addr string
}

func New(addr string) *shell {
	return &shell{
		addr: addr,
	}
}

func (s *shell) Start() error {
	fmt.Printf("usr: ")
	username, err := util.Readline()
	if err != nil {
		panic(err)
	}
	fmt.Printf("pwd: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}
	resp, err := api.Login(string(username), string(password))
	if err != nil {
		log.Error(err)
		return err
	}
	s.ctx = ctx.New(resp.Data)
	err = s.ctx.Connect("tcp", s.addr)
	if err != nil {
		log.Error(err)
		return err
	}
	err = s.ctx.Auth()
	if err != nil {
		log.Error(err)
		return err
	}
	return s.ctx.Listener()
}
