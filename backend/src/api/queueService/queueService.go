package queueService

type QueueService interface {
	Connect() error
	Test() bool
	Send(message string)
}

func Connect() error {

	return nil
}

func Test() bool {
	return true
}
