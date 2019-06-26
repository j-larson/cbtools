package main

import (
	"fmt"
	"os"

	gocb "github.com/couchbase/gocb"
)

func main() {
	actualArgs := os.Args[1:]
	if len(actualArgs) != 4 {
		fmt.Println("Expected arguments: bucket scope collection docid")
		return
	}
	bucketName := actualArgs[0]
	scopeName := actualArgs[1]
	collectionName := actualArgs[2]
	docId := actualArgs[3]
	fmt.Printf("Retrieving from bucket: %s scope: %s collection: %s docid: %s\n", bucketName, scopeName, collectionName, docId)

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
	result, err := collection.Get(docId, &gocb.GetOptions{})
	if err != nil {
		fmt.Printf("Unable to get document %s: %v\n", docId, err)
		return
	}

	var doc interface{}
	err = result.Content(&doc)
	if err != nil {
		fmt.Printf("Unable to parse document: %v\n", err)
		return
	}
	fmt.Printf("The document: %v\n", doc)
}
