package state

import (
	"bytes"
	"fmt"
	"github.com/ratel-online/client/consts"
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/util"
)

type new struct{}

func (*new) Next(ctx *ctx.Context) (consts.StateID, error) {
	options, err := ctx.GetGameTypes()
	if err != nil {
		return 0, err
	}
	buf := bytes.Buffer{}
	buf.WriteString("Please select game type\n")
	for _, option := range options {
		buf.WriteString(fmt.Sprintf("%d.%s\n", option.ID, option.Name))
	}
	fmt.Printf(buf.String())
	gameType, err := util.NextInt()
	if err != nil {
		return 0, err
	}
	room, err := ctx.CreateRoom(gameType)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Create room successful, id : %d\n", room.ID)
	return consts.StateWaiting, nil
}

func (*new) Exit(ctx *ctx.Context) consts.StateID {
	return consts.StateHome
}
