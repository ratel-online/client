package ctx

import (
	"encoding/binary"
	"github.com/ratel-online/client/model"
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/log"
	"github.com/ratel-online/core/network"
	"github.com/ratel-online/core/protocol"
	"github.com/ratel-online/core/util/async"
	"net"
	"testing"
)

func TestContext_Connect(t *testing.T) {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, 4294967294)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "49.235.95.125:9999")
	if err != nil {
		t.Log(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Log(err)
		return
	}
	conn.Write(data)
	select {}
}

func TestContext_Listener(t *testing.T) {
	for i := 1; i <= 5000; i++ {
		id := i
		async.Async(func() {
			_singleTest(int64(id))
		})
	}
	select {}
}

func _singleTest(id int64) {
	ctx := New(model.LoginRespData{
		ID:       id,
		Name:     "a",
		Score:    100,
		Username: "b",
		Token:    "aeiou",
	})
	err := ctx.Connect("tcp", "49.235.95.125:9999")
	if err != nil {
		log.Error(err)
		return
	}
	err = ctx.Auth()
	if err != nil {
		log.Error(err)
		return
	}
	_ = ctx.conn.Accept(func(packet protocol.Packet, conn *network.Conn) {
		data := string(packet.Body)
		if data == consts.IS_START {
			_ = conn.Write(protocol.Packet{
				Body: []byte("2"),
			})
			return
		}
	})
}
