package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Connection URI
// mongodb://<username>:<password>@<host>:<port>/<database-name>
// const uri = "mongodb://mongo:mongo@localhost:27017/todofuken-db"

// mongodb接続用の構造体
type ConnectMongo struct {
	user     string
	password string
	host     string
	port     int
	database string
}

func SetDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (c ConnectMongo) clientConnect() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", c.user, c.password, c.host, c.port, c.database)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client
}
