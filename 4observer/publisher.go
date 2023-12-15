package observer

import "fmt"

type publisher struct {
	subscriberList []Subscriber
}

func GetNewPublisher() Publisher {
	return &publisher{subscriberList: make([]Subscriber, 0)}
}

func (pub *publisher) Register(subs Subscriber) {
	pub.subscriberList = append(pub.subscriberList, subs)
}

func (pub publisher) NotifyAll(msg string) {
	for _, subs := range pub.subscriberList {
		fmt.Println("Publisher notifying Subscriber with ID:", subs.GetId())
		subs.ReactToPublisherMsg(msg)
	}
}
