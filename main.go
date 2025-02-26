package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

type Repo struct {
	Name string `json:"name"`
}

type Event struct {
	Type string `json:"type"`
	Repo Repo `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
}

func readCmd(user *string) string {
	for i, arg := range os.Args {
		if i == 1 {
			*user = arg
		}
	}

	return ""
}

func main() {
	user := ""

	readCmd(&user)

	fURL := fmt.Sprintf("https://api.github.com/users/%v/events", user)
	resp, err := http.Get(fURL)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	} else if resp.StatusCode == 404 {
		log.Fatal("URL not found!")
	}

	data := []Event{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal("Can't parse json in structure!")
	}

	if len(data) == 0 {
		fmt.Println("It's empty here")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', tabwriter.Debug)

	fmt.Fprintf(w, "Event:\tRepo:\tDate:\t\n")

	for _, event := range data {
		fmt.Fprintf(w, "%v\t%v\t%v\t\n", event.Type, event.Repo.Name, event.CreatedAt)
	}

	defer w.Flush()
}