package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://root:root@172.17.0.2:5672/")

	if err != nil {
		fmt.Printf("failed connect to rabbitmq: %e\n", err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	messages, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	// block the main function
	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("successfully connected to rabbitmq instance")
	fmt.Println("[*] - waiting for messages")

	<-forever
}