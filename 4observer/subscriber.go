package observer

import "fmt"

type subscriber struct {
	subscriberId string
}

func GetNewSubscriber(Id string) Subscriber {
	return subscriber{subscriberId: Id}
}

func (s subscriber) ReactToPublisherMsg(msg string) {
	fmt.Printf("%s: received message: %s\n", s.GetId(), msg)
}

func (s subscriber) GetId() string {
	return s.subscriberId
}
