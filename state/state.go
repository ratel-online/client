package state

import (
	"github.com/ratel-online/client/consts"
	"github.com/ratel-online/client/ctx"
)

var states = map[consts.StateID]State{}

func init() {
	register(consts.StateWelcome, &welcome{})
	register(consts.StateHome, &home{})
	register(consts.StateJoin, &join{})
	register(consts.StateNew, &new{})
	register(consts.StateWaiting, &waiting{})
}

func register(id consts.StateID, state State) {
	states[id] = state
}

type State interface {
	Next(ctx *ctx.Context) (consts.StateID, error)
	Exit(ctx *ctx.Context) consts.StateID
}
