package observer

type Publisher interface {
	Register(subscriber Subscriber)
	NotifyAll(msg string)
}

type Subscriber interface {
	ReactToPublisherMsg(msg string)
	GetId() string
}

func Run() {
	pub := GetNewPublisher()

	subs := GetNewSubscriber("1")
	subs1 := GetNewSubscriber("2")
	pub.Register(subs)
	pub.Register(subs1)
	pub.NotifyAll("hi")
}
