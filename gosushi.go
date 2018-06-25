package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	name := "shokunin"
	text := "iccho agari🍣"
	channel := myEnv["SLACK_CHANNEL"]

	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + text + `"}`

	req, err := http.NewRequest(
		"POST",
		myEnv["SLACK_WEBHOOK_URL"],
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(resp)
	defer resp.Body.Close()
}
