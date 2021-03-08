package main

import (
	"fmt"
	"log"
	"strconv"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func init() {

	host.Init()
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("GPIO pins available:\n")
	for _, p := range gpioreg.All() {
		fmt.Printf("- %s: %s\n", p, p.Function())
	}
}

var relays map[string]*relay

type confRelay struct {
	GPIO int    `yaml:"gpio"`
	Name string `yaml:"name"`
}

type relay struct {
	ID    int    `json:"id"`
	GPIO  int    `json:"gpio"`
	Name  string `json:"name"`
	State int    `json:"state"`
}

func (r *relay) getState() int {
	p := gpioreg.ByName(fmt.Sprintf("GPIO%d", r.GPIO))
	if p == nil {
		log.Printf("Failed to find GPIO%d", r.GPIO)
	}

	if p.Read() {
		r.State = 1
	} else {
		r.State = 0
	}

	return r.State
}

func (r *relay) setState(state int) bool {
	r.State = state
	gpioState := gpio.Low

	if r.State == 1 {
		gpioState = gpio.High
	}

	p := gpioreg.ByName(fmt.Sprintf("GPIO%d", r.GPIO))
	if p == nil {
		log.Printf("Failed to find GPIO%d", r.GPIO)
		return false
	}
	if err := p.Out(gpioState); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func loadHardware() {
	relays = make(map[string]*relay)

	for i, r := range conf.Relays {
		stri := strconv.Itoa(i)
		relays[stri] = &relay{
			i,
			r.GPIO,
			r.Name,
			1,
		}
	}
}

func addRelay(gpio int, name string) int {

	relayID := len(relays)

	relays[strconv.Itoa(relayID)] = &relay{
		relayID,
		gpio,
		name,
		0,
	}

	conf.setRelays(relays)

	return relayID
}
