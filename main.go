package main

import (
	"fmt"
	"os"

	"github.com/nmtri1912/blockchain/cmd/blockchaincmd"
)

func main() {
	if err := blockchaincmd.StartCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
