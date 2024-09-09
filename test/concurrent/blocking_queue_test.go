package concurrent

import (
	concurrent "github.com/mobaijaavaer/golang-tools/concurrent/collections"
	"log/slog"
	"math/rand"
	"testing"
	"time"
)

func TestLinkedBlockingQueue_Offer(t *testing.T) {
	queue := concurrent.NewDefaultLinkedBlockingQueue[int]()
	testData := []int{1, 2, 3, 4, 5}

	for _, v := range testData {
		item := &v
		success := queue.Offer(item)
		if !success {
			t.Errorf("Offer failed for item: %d", v)
		}
	}
}

func TestLinkedBlockingQueue_Poll(t *testing.T) {
	queue := concurrent.NewDefaultLinkedBlockingQueue[int]()
	testData := []int{1, 2, 3, 4, 5}

	for _, v := range testData {
		item := &v
		queue.Offer(item)
	}

	for _, v := range testData {
		polledItem := queue.Poll()
		if polledItem == nil || *polledItem != v {
			t.Errorf("Polled item is not as expected. Got: %d, want: %d", *polledItem, v)
		}
	}
}

func TestProducerAndConsumer(t *testing.T) {
	queue := concurrent.NewDefaultLinkedBlockingQueue[int]()
	go func() {

		for {
			time.Sleep(1 * time.Second)
			r := rand.Int()
			queue.Offer(&r)

		}
	}()

	go func() {
		for {
			value := queue.Poll()
			slog.Info("value", "value", *value)
		}

	}()

	ch := make(chan struct{}, 1)
	<-ch

}
