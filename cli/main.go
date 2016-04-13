package main

import (
	"github.com/rusnassonov/go-todoist"
	"log"
	"os"
	"errors"
)

func readFromFile(name string) (string , error) {
	file, err := os.Open(name)

	if err != nil {
		return "", err
	}

	var data []byte
	_, err = file.Read(data)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func readFromEnv(name string) (string, error) {
	token := os.Getenv(name)

	if token == "" {
		return "", errors.New("Token not found")
	}

	return token, nil
}

func ReadToken(name string) (string, error) {
	token, err := readFromFile(name)

	if err == nil {
		return token, nil
	}

	token, err = readFromEnv(name)

	if err == nil {
		return token, nil
	}

	return "", err
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	token, err := ReadToken("todoist_api_token")

	if err != nil {
		log.Print(err)
		return
	}

	api := go_todoist.NewTodoistAPI(token)
	api.Sync()

	if err != nil {
		log.Print(err)
	}


	log.Println(api.User)
	//log.Println(api.User.Token)
	//log.Println(api.SeqNoGlobal)

}
