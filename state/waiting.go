package state

import (
	"bytes"
	"fmt"
	"github.com/ratel-online/client/consts"
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/util"
	"github.com/ratel-online/core/errors"
	"strings"
	"time"
)

type waiting struct{}

func (s *waiting) Next(ctx *ctx.Context) (consts.StateID, error) {
	room := service.GetRoom(player.RoomID)
	if room == nil {
		return 0, consts.ErrorsExist
	}
	access, err := waitingForStart(player, room)
	if err != nil {
		return 0, err
	}
	if access {
		if room.Type == consts.GameTypeClassic {
			return consts.StateClassics, nil
		} else if room.Type == consts.GameTypeLaiZi {
			return consts.StateLaiZi, nil
		}
	}
	return s.Exit(player), nil
}

func (*waiting) Exit(ctx *ctx.Context) consts.StateID {

	room := service.GetRoom(player.RoomID)
	if room != nil {
		isOwner := room.Creator == player.ID
		service.LeaveRoom(room.ID, player.ID)
		service.Broadcast(room.ID, fmt.Sprintf("%s exited room! room current has %d players\n", player.Name, room.Players))
		if isOwner {
			newOwner := service.GetPlayer(room.Creator)
			service.Broadcast(room.ID, fmt.Sprintf("%s become new owner\n", newOwner.Name))
		}
	}
	return consts.StateHome
}

func waitingForStart(player *service.Player, room *service.Room) (bool, error) {
	access := false
	for {
		signal, err := util.NextString(time.Second)
		if err != nil && err != errors.Timeout {
			return access, err
		}
		if room.State == consts.RoomStateRunning {
			access = true
			break
		}
		signal = strings.ToLower(signal)
		if signal == "ls" || signal == "v" {
			viewRoomPlayers(room, player)
		} else if (signal == "start" || signal == "s") && room.Creator == player.ID && room.Players > 1 {
			access = true
			room.Lock()
			room.Game, err = initGame(room)
			if err != nil {
				return access, err
			}
			room.State = consts.RoomStateRunning
			room.Unlock()
			break
		}
	}
	return access, nil
}

func viewRoomPlayers(room *service.Room, currPlayer *service.Player) {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%-20s%-10s%-10s\n", "Name", "Score", "Title"))
	for playerId := range service.GetRoomPlayers(room.ID) {
		title := "player"
		if playerId == room.Creator {
			title = "owner"
		}
		player := service.GetPlayer(playerId)
		buf.WriteString(fmt.Sprintf("%-20s%-10d%-10s\n", player.Name, player.Score, title))
	}
	_ = currPlayer.WriteString(buf.String())
}

func initGame(room *service.Room) (*service.Game, error) {
	if room.Type == consts.GameTypeClassic {
		return classics.InitGame(room)
	} else if room.Type == consts.GameTypeLaiZi {
		return laizi.InitGame(room)
	}
	return nil, nil
}
