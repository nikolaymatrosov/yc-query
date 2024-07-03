package main

import (
	"github.com/nikolaymatrosov/yc-query/pkg/cmd"
)

func main() {
	err := cmd.Root().Execute()
	if err != nil {
		print(err)
		return
	}
}
