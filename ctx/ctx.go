package ctx

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ratel-online/client/model"
	"github.com/ratel-online/client/util"
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/log"
	modelx "github.com/ratel-online/core/model"
	"github.com/ratel-online/core/network"
	"github.com/ratel-online/core/protocol"
	"github.com/ratel-online/core/util/async"
	"net"
	"net/url"
	"strings"
)

type Context struct {
	id     int64
	name   string
	score  int64
	token  string
	roomId int64

	conn *network.Conn
}

type netConnector func(addr string) (*network.Conn, error)

var netConnectors = map[string]netConnector{
	"tcp": tcpConnect,
	"ws":  websocketConnect,
}

func New(user model.LoginRespData) *Context {
	return &Context{
		id:    user.ID,
		name:  user.Name,
		score: user.Score,
		token: user.Token,
	}
}

func (c *Context) Connect(net string, addr string) error {
	if connector, ok := netConnectors[net]; ok {
		conn, err := connector(addr)
		if err != nil {
			return err
		}
		c.conn = conn
		return nil
	}
	return errors.New(fmt.Sprintf("unsupported net type: %s", net))
}

func (c *Context) Auth() error {
	return c.conn.Write(protocol.ObjectPacket(modelx.AuthInfo{
		ID:    c.id,
		Name:  c.name,
		Score: c.score,
	}))
}

func (c *Context) Listener() error {
	is := false
	async.Async(func() {
		for {
			line, err := util.Readline()
			if err != nil {
				log.Panic(err)
			}
			if !is {
				continue
			}
			err = c.conn.Write(protocol.Packet{
				Body: line,
			})
			if err != nil {
				continue
			}
			is = false
		}
	})
	return c.conn.Accept(func(packet protocol.Packet, conn *network.Conn) {
		data := string(packet.Body)
		if data == consts.IS {
			if !is {
				fmt.Print(fmt.Sprintf("\r\r[%s@ratel %s]# ", strings.TrimSpace(strings.ToLower(c.name)), "~"))
			}
			is = true
			return
		}
		if is {
			fmt.Print("\r\r" + data + fmt.Sprintf("\r\r[%s@ratel %s]# ", strings.TrimSpace(strings.ToLower(c.name)), "~"))
		} else {
			fmt.Print(data)
		}
	})
}

func tcpConnect(addr string) (*network.Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("tcp server error: %v", err))
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("tcp server error: %v", err))
	}
	return network.Wrapper(protocol.NewTcpReadWriteCloser(conn)), nil
}

func websocketConnect(addr string) (*network.Conn, error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ws server error: %v", err))
	}
	return network.Wrapper(protocol.NewWebsocketReadWriteCloser(conn)), nil
}
