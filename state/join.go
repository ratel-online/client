package state

import (
	"bytes"
	"fmt"
	"github.com/ratel-online/client/consts"
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/util"
	"github.com/ratel-online/core/log"
)

type join struct{}

func (s *join) Next(ctx *ctx.Context) (consts.StateID, error) {
	buf := bytes.Buffer{}
	rooms, err := ctx.GetRooms()
	if err != nil {
		log.Error(err)
		return 0, err
	}
	buf.WriteString(fmt.Sprintf("%-10s%-10s%-10s%-10s\n", "ID", "Type", "Players", "State"))
	for _, room := range rooms {
		buf.WriteString(fmt.Sprintf("%-10d%-10s%-10d%-10s\n", room.ID, room.TypeDesc, room.Players, room.StateDesc))
	}
	fmt.Printf(buf.String())
	roomId, err := util.NextInt64()
	if err != nil {
		return 0, err
	}
	err = ctx.JoinRoom(roomId)
	if err != nil {
		return 0, err
	}
	return consts.StateWaiting, nil
}

func (*join) Exit(ctx *ctx.Context) consts.StateID {
	return consts.StateHome
}
