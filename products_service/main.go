package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

const (
	productsTopic = "product_topic"
)

type UserEvent struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func main() {
	logger := watermill.NewStdLogger(false, false)
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	ctx := context.Background()

	messages, err := pubSub.Subscribe(ctx, productsTopic)
	if err != nil {
		log.Fatalf("could not subscribe to topic: %v", err)
	}

	go func() {
		for msg := range messages {
			log.Printf("Product Service received user event: %s", string(msg.Payload))
			msg.Ack()
		}
	}()

	http.HandleFunc("/products/create_product", func(w http.ResponseWriter, r *http.Request) {
		userEvent := UserEvent{
			UserID:   "USERID1",
			Username: "USERNAME1",
		}

		msg := message.NewMessage(watermill.NewUUID(), []byte(userEvent.UserID))
		if err := pubSub.Publish(productsTopic, msg); err != nil {
			http.Error(w, "Failed to publish message", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created and event published"))
	})

	log.Println("Starting Product Service on port 8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Product Service failed: %v", err)
	}
}
