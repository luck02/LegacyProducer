package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
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
	for index, element := range legacyEvents.LegacyEvents {
		fmt.Print(index)
		fmt.Print(element)
		fmt.Println()
	}
}
