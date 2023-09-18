package main

import (
	"fmt"
	"net/http"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collega struct {
	ID       interface{} `bson:"_id" json:"_id"`
	Nome     string      `bson:"nome" json:"nome"`
	Cognome  string      `bson:"cognome" json:"cognome"`
	Mansione string      `bson:"mansione" json:"mansione"`
}

func main() {
	// Crea un handler che catturi tutte le chiamate al percorso /example
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		opts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb")
		// Create a new client and connect to the server
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect(context.TODO())

		var collega Collega

		col := client.Database("data").Collection("colleghi")
		err = col.FindOne(context.TODO(), bson.D{{"nome", "Mirko"}}).Decode(&collega)
		if err != nil {
			panic(err)
		}

		fmt.Println(collega)

		// Scrivi il JSON nella risposta
		fmt.Fprintf(w, "%s", collega)

	})

	// Avvia il server HTTP
	http.ListenAndServe(":8080", handler)
}
