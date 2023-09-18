package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Create a client
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the client
	defer client.Disconnect(context.Background())

	// Print the version
	fmt.Println(client.Version())
}
