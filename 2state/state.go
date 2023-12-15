package state

import (
	"errors"
	"fmt"
)

var (
	TVOn  = on{}
	TVOff = off{}
)

type tvState interface {
	powerOn() error
	powerOff() error
	watch(channel string) error
}

type on struct{}

func (o on) powerOn() error {
	return errors.New("TV already on")
}

func (o on) powerOff() error {
	fmt.Println("Turning Off TV")
	return nil
}

func (o on) watch(channel string) error {
	fmt.Println("watching " + channel)
	return nil
}

type off struct{}

func (o off) powerOn() error {
	fmt.Println("Turning On TV")
	return nil
}

func (o off) powerOff() error {
	return errors.New("TV already off")
}

func (o off) watch(channel string) error {
	return errors.New("cant watch tv is off")
}
