/**
2 * @Author: Nico
3 * @Date: 2021/5/23 20:28
4 */
package util

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "strings"
)

var (
    in     *bufio.Reader = bufio.NewReader(os.Stdin)
    out    *bufio.Writer = bufio.NewWriter(os.Stdout)
    buffer               = bytes.Buffer{}
)

func Readline() ([]byte, error) {
    lines, err := in.ReadBytes('\n')
    if err != nil {
        return nil, err
    }
    return []byte(strings.TrimSpace(string(lines[0 : len(lines)-1]))), nil
}

func ReadBuffer() []byte {
    return buffer.Bytes()
}

func EraseLine() {
    fmt.Printf("\033[1A")
    fmt.Printf("\r\r")
}
