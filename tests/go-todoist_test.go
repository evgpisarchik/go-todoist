package tests

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/rusnassonov/go-todoist"
)

var api *go_todoist.TodoistAPI

func readFromFile(name string) (string, error) {
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
func init() {
	token, err := ReadToken("todoist_api_token")

	if err != nil {
		log.Println("Todoist API token not found")
	}
	api = go_todoist.NewTodoistAPI(token)
}

func TestAddProject(t *testing.T) {
	err := api.Sync()

	if err != nil {
		t.Fail()
	}

	p := go_todoist.Project{Name: "Hello world!"}
	api.Projects.Add(&p)

	err = api.Sync()

	if err != nil {
		t.Fatalf("Add project error: %v", err)
	}

	if p.ID == "0" {
		t.Errorf("Project ID is 0")
	}

}

func TestDeleteProject(t *testing.T) {
	err := api.Sync()

	if err != nil {
		t.Fail()
	}

	p := go_todoist.Project{Name: "Hello world!"}
	api.Projects.Add(&p)

	api.Sync()

	api.Projects.Delete(&p)

	err = api.Sync()

	if err != nil {
		t.Errorf("Fail delete project: %v", err)
	}
}

func TestDeleteAll(t *testing.T) {
	api.Sync()

	api.Projects.DeleteAll()

	api.Sync()
}
