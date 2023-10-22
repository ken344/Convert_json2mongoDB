package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type mongoParams struct {
	host       string
	user       string
	password   string
	database   string
	collection string
}

func newMongoParams(host string, user string, password string, database string, collection string) *mongoParams {
	mg := new(mongoParams)
	mg.host = host
	mg.user = user
	mg.password = password
	mg.database = database
	mg.collection = collection
	return mg
}

// jsonファイルをmongoDBにインポートする
func (m mongoParams) importJson(filePath string) {
	// mongoimportを使用するためには、mongodb-database-toolsをインストールする必要がある。
	//https://www.mongodb.com/docs/database-tools/installation/installation-macos/
	cmd := exec.Command("/usr/local/bin/mongoimport", "-h", m.host, "-u", m.user, "-p", m.password, "--db", m.database, "--collection", m.collection, "--file", filePath, "--jsonArray")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}

// .envファイルを読み込む
func SetDotenv(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ファイルのパスを取得して配列にする。
func getFilePaths(dirPath string, extensionName string) []string {
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

		// 拡張子がにextensionNameであった場合は配列に格納する
		if strings.EqualFold(filepath.Ext(path), "."+extensionName) {
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

	// .envファイルを読み込む/
	SetDotenv("../../.env")
	//import用の構造体を作成する
	mongoImport := newMongoParams(os.Getenv("MONGO_HOST"), os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_COLLECTION"))

	// 指定したディレクトリ内に存在するファイルから、指定した拡張子のファイルのパスを配列に格納する
	targetDirPath := "../../input_data/"
	choiceExtensionName := "json"
	jsonFilePath := getFilePaths(targetDirPath, choiceExtensionName)

	// jsonファイルをmongoDBにインポートする
	for _, filePath := range jsonFilePath {
		mongoImport.importJson(filePath)
	}
}
