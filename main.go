package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	TIMEDURATION = 5
)

func main() {
	server := scanServer()

	client := &http.Client{}

	for {
		fmt.Println()
		fmt.Println("Sending a request...")

		responseCode := sendRequest(server, client)
		if responseCode == 200 {
			fmt.Println("The server is available!")
		} else {
			fmt.Println("Error detected! The server is down!")
		}

		fmt.Println()
		fmt.Println("Repeat the request through", TIMEDURATION, "sec.")
		time.Sleep(time.Second * TIMEDURATION)
	}
}

func sendRequest(server string, client *http.Client) int {
	resp, err := client.Get(server)
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode
}

func scanServer() string {
	var server string

	_, err := fmt.Scan(&server)
	if err != nil {
		log.Fatal(err)
	}

	return server
}
