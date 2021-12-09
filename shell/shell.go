package shell

import (
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/model"
	"github.com/ratel-online/client/util"
	"github.com/ratel-online/core/log"
	"os"
	"time"
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
	//fmt.Printf("usr: ")
	//username, err := util.Readline()
	//if err != nil {
	//    panic(err)
	//}
	//fmt.Printf("pwd: ")
	//password, err := gopass.GetPasswd()
	//if err != nil {
	//    panic(err)
	//}
	//resp, err := api.Login(string(username), string(password))
	//if err != nil {
	//    log.Error(err)
	//    return err
	//}
	name := util.RandomName()
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	//s.addr = "49.235.95.125:9999"
	//if len(os.Args) > 2 {
	//	s.addr = os.Args[2]
	//}
	s.ctx = ctx.New(model.LoginRespData{
		ID:       time.Now().UnixNano(),
		Name:     name,
		Score:    100,
		Username: name,
		Token:    "aeiou",
	})
	err := s.ctx.Connect("tcp", s.addr)
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
