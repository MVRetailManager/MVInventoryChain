package main

import (
	"os"

	"github.com/MVRetailManager/MVInventoryChain/cli"
	"github.com/MVRetailManager/MVInventoryChain/logging"
)

func init() {
	logging.SetupLogger()

	logging.InfoLogger.Println("Starting MVInventoryChain...")
}

func main() {
	defer os.Exit(0)
	cmdline := cli.CLI{}
	cmdline.Run()
}
