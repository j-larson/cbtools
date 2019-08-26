package main

import (
	"fmt"
	"os"

	"encoding/json"
	gocb "github.com/couchbase/gocb"
)

func main() {
	actualArgs := os.Args[1:]
	if len(actualArgs) != 5 {
		fmt.Println("Expected arguments: bucket scope collection docid json_text")
		return
	}
	bucketName := actualArgs[0]
	scopeName := actualArgs[1]
	collectionName := actualArgs[2]
	docId := actualArgs[3]
	jsonText := actualArgs[4]
	fmt.Printf("Inserting to bucket: %s scope: %s collection: %s docid: %s\n", bucketName, scopeName, collectionName, docId)

	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			"Administrator",
			"password",
		},
	}
	cluster, err := gocb.Connect("localhost", opts)
	if err != nil {
		fmt.Printf("Unable to connect: %v\n", err)
		return
	}
	bucket := cluster.Bucket(bucketName, &gocb.BucketOptions{})
	scope := bucket.Scope(scopeName)
	collection := scope.Collection(collectionName, &gocb.CollectionOptions{})
	message := json.RawMessage(jsonText)
	_, err = collection.Insert(docId, message, &gocb.InsertOptions{})
	if err != nil {
		fmt.Printf("Unable to insert document %s: %v\n", docId, err)
		return
	}
	fmt.Printf("Successfully inserted document %s with body %s\n", docId, jsonText)
}
