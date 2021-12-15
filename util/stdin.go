package util

import (
	"bufio"
	"github.com/ratel-online/core/errors"
	"github.com/ratel-online/core/util/async"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	in = bufio.NewReader(os.Stdin)
	ch = make(chan *string)
)

func init() {
	async.Async(func() {
		for {
			line, err := readline()
			if err != nil {
				break
			}
			ch <- &line
		}
	})
}

func readline() (string, error) {
	lines, err := in.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(lines[0 : len(lines)-1])), nil
}

func NextString(timeout ...time.Duration) (string, error) {
	var line *string
	if len(timeout) > 0 {
		select {
		case line = <-ch:
		case <-time.After(timeout[0]):
			return "", errors.Timeout
		}
	} else {
		line = <-ch
	}
	if line == nil {
		return "", errors.ChanClosed
	}
	return *line, nil
}

func NextInt(timeout ...time.Duration) (int, error) {
	v, err := NextInt64(timeout...)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func NextInt64(timeout ...time.Duration) (int64, error) {
	line, err := NextString(timeout...)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(line, 10, 64)
}
