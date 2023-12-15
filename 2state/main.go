package state

import (
	"fmt"
)

type TV struct {
	currentTvState tvState
}

func NewTV() *TV {
	return &TV{
		currentTvState: TVOff,
	}
}

func (sc *TV) PowerOn() {
	err := sc.currentTvState.powerOn()
	if err != nil {
		fmt.Printf("error turning on tv: %s\n", err.Error())
		return
	}
	sc.currentTvState = TVOn
}

func (sc *TV) PowerOff() {
	err := sc.currentTvState.powerOff()
	if err != nil {
		fmt.Printf("error turning off tv: %s\n", err.Error())
		return
	}
	sc.currentTvState = TVOff
}

func (sc *TV) Watch(channel string) {
	err := sc.currentTvState.watch(channel)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}

func Run() {
	tvContext := NewTV()
	tvContext.Watch("net tv")
	tvContext.PowerOff()
	tvContext.PowerOn()
	tvContext.Watch("net tv")
	tvContext.PowerOn()
	tvContext.PowerOff()
}
