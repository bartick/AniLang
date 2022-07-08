package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bartick/AniLang/utils"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	utils.AddFlag("m", "To set value xyz")
	utils.AddBoolFlag("ml", "To add math library")
	utils.FlagParser()

	if utils.FilePath == "" {
		utils.TerminalReader()
	} else {
		_, _ = utils.FileReader(utils.FilePath)
	}

}
