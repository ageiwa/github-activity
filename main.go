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

func main() {
	fURL := fmt.Sprintf("https://api.github.com/users/%v/events", os.Args[1])
	resp, err := http.Get(fURL)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	data := []Event{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err.Error())
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