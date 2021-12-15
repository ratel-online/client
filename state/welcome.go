package state

import (
	"fmt"
	"github.com/ratel-online/client/consts"
	"github.com/ratel-online/client/ctx"
)

type welcome struct{}

func (*welcome) Next(ctx *ctx.Context) (consts.StateID, error) {
	fmt.Printf("Hi %s, Welcome to ratel online! \n", ctx.Name)
	return consts.StateHome, nil
}

func (*welcome) Exit(ctx *ctx.Context) consts.StateID {
	return 0
}
