package pubsub

type Subscriber struct{}

type SubscribeServer struct {
}

func NewSubscriber() *SubscribeServer {
	return &SubscribeServer{}
}
