package shell

import (
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/model"
	"github.com/ratel-online/client/util"
	"github.com/ratel-online/core/log"
	"time"
)

type shell struct {
	ctx  *ctx.Context
	addr string
	name string
}

func New(addr, name string) *shell {
	return &shell{
		addr: addr,
		name: name,
	}
}

func (s *shell) Start() error {
	name := util.RandomName()
	if s.name != "" {
		name = s.name
	}
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
