package main

import (
	"fmt"
	"os"
)

func main() {
	filePaths := getFilePaths(sourceDataDirPath)
	for _, filePath := range filePaths {
		println(filePath)
	}

	// .envファイルの読み込み
	SetDotenv()

	mongoClient := ConnectMongo{
		user:     os.Getenv("MONGO_USER"),
		password: os.Getenv("MONGO_PASSWORD"),
		host:     "localhost",
		port:     27017,
		database: os.Getenv("MONGO_DATABASE"),
	}

	_ = mongoClient.clientConnect()

	fmt.Println("Successfully connected and pinged.")

}
