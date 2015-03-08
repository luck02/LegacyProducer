package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type LegacyEvents struct {
	LegacyEvents []LegacyEvent
}
type LegacyEvent struct {
	WorkflowId string
	Type       string
	KeyType    string
	KeyValue   int
}

func InitLogger() {
	logFile, err := os.OpenFile("sarama.debug.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	sarama.Logger = log.New(logFile,
		"PREFIX: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InitLogger()
	file, e := ioutil.ReadFile("./TestInput.json")
	if e != nil {
		fmt.Printf("FIle error: %v\n", e)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(file))
	var legacyEvents LegacyEvents

	e = json.Unmarshal(file, &legacyEvents)
	if e != nil {
		fmt.Printf("error: ", e)
		os.Exit(1)
	}

	fmt.Printf("Results: %v\n", legacyEvents)
	PublishEvents(legacyEvents)
}

func PublishEvents(legacyEvents LegacyEvents) {
	producer, err := sarama.NewProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	for index, element := range legacyEvents.LegacyEvents {
		fmt.Print(index)
		fmt.Print(element)
		fmt.Println()

		elementBytes, err := json.Marshal(element)
		if err != nil {
			fmt.Println(err)
		}

		elementString := elementBytes[:]
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: "LegacyEvents", Key: nil, Value: sarama.StringEncoder(elementString)}:
			fmt.Println(">Legacy Events QUeued")
		case err := <-producer.Errors():
			panic(err.Err)

		}

	}
}
