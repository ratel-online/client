package ctx

import (
	"errors"
	"fmt"
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/model"
	"strconv"
)

func (c *Context) GetRooms() ([]model.Room, error) {
	resp, err := c.Request(consts.Service, consts.ServiceGetRooms, nil)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(fmt.Sprintf("err code: %v, msg: %v", resp.Code, resp.Msg))
	}
	rooms := make([]model.Room, 0)
	return rooms, resp.Unmarshal(&rooms)
}

func (c *Context) JoinRoom(roomId int64) error {
	resp, err := c.Request(consts.Service, consts.ServiceJoinRoom, strconv.FormatInt(roomId, 10))
	if err != nil {
		return err
	}
	if resp.Code != 0 {

		return errors.New(fmt.Sprintf("err code: %v, msg: %v", resp.Code, resp.Msg))
	}
	c.RoomId = roomId
	return nil
}

func (c *Context) CreateRoom(gameType int) (*model.Room, error) {
	resp, err := c.Request(consts.Service, consts.ServiceCreateRoom, strconv.Itoa(gameType))
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(fmt.Sprintf("err code: %v, msg: %v", resp.Code, resp.Msg))
	}
	room := &model.Room{}
	err = resp.Unmarshal(room)
	if err != nil {
		return nil, err
	}
	c.RoomId = room.ID
	return room, nil
}

func (c *Context) GetGameTypes() ([]model.Option, error) {
	resp, err := c.Request(consts.Service, consts.ServiceGetGameTypes, nil)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(fmt.Sprintf("err code: %v, msg: %v", resp.Code, resp.Msg))
	}
	options := make([]model.Option, 0)
	return options, resp.Unmarshal(&options)
}
