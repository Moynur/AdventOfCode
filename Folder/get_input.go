package AoC

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)
var sessionID = ""

func getTextFile(dayNeeded, yearNeeded int) {
	sessionID := os.Getenv("session")
	if sessionID == "" {
		log.Println("You need to export your cookie")
	}
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input",yearNeeded, dayNeeded)
	loginSession := http.Cookie{
		Name:  "session",
		Value: sessionID,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.AddCookie(&loginSession)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		log.Println(err)
	}
	defer resp.Body.Close()
	out, err := os.Create("input.txt")
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}

func GetInput(dayNeeded, year int) []string {
	getTextFile(dayNeeded, year)
	var content []string
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return content
}

func GetInputToday() []string{
	getTextFile(time.Now().Day(), time.Now().Year())
	var content []string
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return content
}