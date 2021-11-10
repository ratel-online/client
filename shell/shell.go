package shell

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/render"
	"net/http"
	"strings"
)

var servers = []string{
	"https://raw.githubusercontent.com/awesome-cmd/chat/main/servers.json",
	"https://gitee.com/ainilili/chat/raw/main/servers.json",
}

type shell struct {
	ctx      *ctx.Context
	position *directory
	addr    string
}

func New(addr string) *shell {
	return &shell{
		ctx: ctx.NewContext(),
		position: root,
		addr:    addr,
	}
}

func (s *shell) Start() {
	for {
		fmt.Printf("[%s@hxd %s]# ", s.ctx.Name, s.position.name)
		inputs, err := render.Readline()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if strings.TrimSpace(string(inputs)) != "" {
			res, err := s.position.action(s, inputs)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Print(res)
		}
	}
}

func (s *shell) refreshServerList() error {
	serverList := make([]string, 0)
	if s.addrs != "" {
		serverList = append(serverList, s.addrs+"|default")
	} else {
		for _, server := range servers {
			resp, err := http.Get(server)
			if err != nil && resp == "" {
				continue
			}
			json.Unmarshal([]byte(resp), &serverList)
		}
	}
	if len(serverList) == 0 {
		return errors.New("no available server. ")
	}
	s.position.reset()
	for _, v := range serverList {
		serverInfo := strings.Split(v, "|")
		s.position.add(newDirectory(strings.ToLower(serverInfo[1]), strings.ToLower(serverInfo[0]), serverActions, baseActions))
	}
	return nil
}
