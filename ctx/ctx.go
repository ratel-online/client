package ctx

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ratel-online/client/api"
	"github.com/ratel-online/core/consts"
	errorx "github.com/ratel-online/core/errors"
	"github.com/ratel-online/core/model"
	"github.com/ratel-online/core/network"
	"github.com/ratel-online/core/protocol"
	"github.com/ratel-online/core/util/json"
	"net"
	"net/url"
	"sync"
	"time"
)

const cleanLine = "\r\r                                                                                              \r\r"

type Context struct {
	sync.Mutex
	ID     int64
	Name   string
	Score  int64
	Token  string
	RoomId int64

	conn     *network.Conn
	channels map[int]chan *model.Resp
}

type netConnector func(addr string) (*network.Conn, error)

var netConnectors = map[string]netConnector{
	"tcp": tcpConnect,
	"ws":  websocketConnect,
}

func New(user api.LoginRespData) *Context {
	channels := map[int]chan *model.Resp{}
	for i := 1; i <= 3; i++ {
		channels[i] = make(chan *model.Resp, 10)
	}
	return &Context{
		ID:       user.ID,
		Name:     user.Name,
		Score:    user.Score,
		Token:    user.Token,
		channels: channels,
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
	return c.conn.Write(protocol.ObjectPacket(model.AuthInfo{
		ID:    c.ID,
		Name:  c.Name,
		Score: c.Score,
		Token: c.Token,
	}))
}

func (c *Context) Listener() error {
	return c.conn.Accept(func(packet protocol.Packet, conn *network.Conn) {
		resp := &model.Resp{}
		_ = json.Unmarshal(packet.Body, resp)
		if channel, ok := c.channels[resp.Type]; ok {
			channel <- resp
		}
	})
}

func (c *Context) print(str string) {
	c.Lock()
	defer c.Unlock()
	fmt.Print(str)
}

func (c *Context) Request(t int, code int, data interface{}) (*model.Resp, error) {
	err := c.conn.Write(protocol.Packet{
		Body: json.Marshal(model.Req{
			Type: t,
			Code: code,
			Data: json.Marshal(data),
		}),
	})
	if err != nil {
		return nil, err
	}
	if t == consts.Service {
		select {
		case resp := <-c.channels[t]:
			return resp, nil
		case <-time.After(3 * time.Second):
			return nil, errorx.Timeout
		}
	}
	return nil, nil
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
