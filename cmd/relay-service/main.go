package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	setupCloseHandler()
	conf = getConfig()

	loadHardware()

	setupGinRouter()
	runGinRouter()

}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Shutting down")
		conf.write()
		os.Exit(0)
	}()
}
