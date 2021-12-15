package state

import (
	"bytes"
	"fmt"
	"github.com/ratel-online/client/consts"
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/util"
	"github.com/ratel-online/core/errors"
)

type home struct{}

func (*home) Next(ctx *ctx.Context) (consts.StateID, error) {
	buf := bytes.Buffer{}
	buf.WriteString("1.Join\n")
	buf.WriteString("2.New\n")
	fmt.Printf(buf.String())

	selected, err := util.NextInt()
	if err != nil {
		return 0, errors.InputInvalid
	}
	if selected == 1 {
		return consts.StateJoin, nil
	} else if selected == 2 {
		return consts.StateNew, nil
	}
	return 0, errors.InputInvalid
}

func (*home) Exit(ctx *ctx.Context) consts.StateID {
	return 0
}
