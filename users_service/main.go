package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

const (
	usersTopic = "users_topic"
)

func main() {
	port := flag.Int("port", 8081, "Port to run the User Service on")
	flag.Parse()

	logger := watermill.NewStdLogger(false, false)
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	ctx := context.Background()

	messages, err := pubSub.Subscribe(ctx, usersTopic)
	if err != nil {
		log.Fatalf("could not subscribe to topic: %v", err)
	}

	go func() {
		for msg := range messages {
			log.Printf("User Service received user event: %s", string(msg.Payload))
			msg.Ack()
		}
	}()

	http.HandleFunc("/users/create_user", func(w http.ResponseWriter, r *http.Request) {
		userID := "user-123"

		msg := message.NewMessage(watermill.NewUUID(), []byte(userID))
		if err := pubSub.Publish(usersTopic, msg); err != nil {
			log.Printf("could not publish message: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("User event published: %s", userID)
		w.Write([]byte("User created and event published"))
	})

	address := fmt.Sprintf(":%d", *port)
	log.Printf("Starting User Service on port %d", *port)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("User Service failed: %v", err)
	}
}
