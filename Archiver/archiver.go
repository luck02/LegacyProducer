package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	master, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> master consumer ready")
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	consumer, err := master.ConsumePartition("LegacyEvents", 0, 0)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("> Consumer Ready")
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	msgCount := 0
	legacyEventChannel := make(chan string)
	go WriteArchiveLog(legacyEventChannel)

consumerLoop:
	for {
		select {
		case err := <-consumer.Errors():
			panic(err)
		case mesg := <-consumer.Messages():
			msgCount++
			fmt.Println(mesg)
			legacyEventChannel <- mesg
		case <-time.After(5 * time.Second):
			fmt.Println("> timed out")
			break consumerLoop
		}
	}
	fmt.Println("Got: ", msgCount, " messages.")

}

func WriteArchiveLog(legacyEventChannel chan string) {
	for {
		select {
		case msg := <-legacyEventChannel:
			FormatAndArchiveMessage(msg)
		default:
		}
	}
}

func FormatAndArchiveMessage(messageString string) {
	file, err := os.Create("./archive")
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(messageString)

	if err != nil {
		panic(err)
	}
}
