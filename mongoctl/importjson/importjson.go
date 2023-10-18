package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type mongoParams struct {
	host       string
	user       string
	password   string
	database   string
	collection string
}

func (m mongoParams) importJson(filePath string) {
	cmd := exec.Command("mongoimport", "-u", m.user, "-p", m.password, "--db", m.database, "--collection", m.collection, "--file", filePath, "--jsonArray")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}

func SetDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ファイルのパスを取得して配列にする。
func getFilePaths(dirPath string, keyword string) []string {
	// ファイルのパスを格納する配列
	var filePaths []string

	// ディレクトリの中身を読み込む
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error: path %v, err %v\n", path, err)
		}

		// ディレクトリは無視する
		if info.IsDir() {
			return nil
		}

		// ファイル名にkeywordが含まれている場合は配列に格納する
		if filepath.Base(path) == keyword {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", dirPath, err)
	}

	return filePaths
}

func main() {

	SetDotenv()
	mongoClient := mongoParams{
		host:       os.Getenv("MONGO_HOST"),
		user:       os.Getenv("MONGO_USER"),
		password:   os.Getenv("MONGO_PASSWORD"),
		database:   os.Getenv("MONGO_DATABASE"),
		collection: os.Getenv("MONGO_COLLECTION"),
	}

	targetDirPath := "../../input_data/"
	choiceKey := "json"
	jsonFilePath := getFilePaths(targetDirPath, choiceKey)

	for _, filePath := range jsonFilePath {
		mongoClient.importJson(filePath)
	}
}
