// Copyright 2018 Kuei-chun Chen. All rights reserved.

package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/simagix/mongo-go-examples/examples"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func main() {
	var err error
	var client *mongo.Client
	var connStr connstring.ConnString
	collection := flag.String("collection", "", "collection to watch")
	pipe := flag.String("pipeline", "", "aggregation pipeline")

	flag.Parse()
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if client, err = examples.GetMongoClient(flag.Arg(0)); err != nil {
		panic(err)
	}

	var pipeline = []bson.D{}
	if *pipe != "" {
		pipeline = examples.MongoPipeline(*pipe)
	}

	stream := examples.NewChangeStream()
	stream.SetCollection(*collection)
	stream.SetDatabase(connStr.Database)
	stream.SetPipeline(pipeline)
	stream.Watch(client, echo)
}

func echo(doc bson.M) {
	var b []byte
	b, _ = json.MarshalIndent(doc, "", "  ")
	fmt.Println(string(b))
}
