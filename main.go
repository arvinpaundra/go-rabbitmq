package main

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	// create connection to rabbitmq
	conn, err := amqp091.Dial("amqp://root:root@172.17.0.2:5672/")

	if err != nil {
		fmt.Printf("failed connect to rabbit: %e\n", err)
		panic(err)
	}

	// close rabbitmq when application stop
	defer conn.Close()

	fmt.Println("successfully connected to rabbitmq instance")

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(queue)

	ctx := context.Background()

	err = ch.PublishWithContext(
		ctx,
		"",
		"TestQueue",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body: []byte("Hello World"),
		},
	)

	fmt.Println("successfully published message to queue")
}