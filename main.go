package main

import (
	"fmt"
	"os"

	"github.com/ratel-online/client/shell"
	"github.com/ratel-online/core/log"
)

func main() {
	if len(os.Args) > 0 {
        for _, arg := range os.Args {
            if arg=="-help" {
				fmt.Print(HELP)
				os.Exit(0)
			}
        }
    }
	log.Error(shell.New("127.0.0.1:9999").Start())
}
