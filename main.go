package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const omdbURL = "http://www.omdbapi.com/"
const omdbKey = "298f93b0"
const testUrl = "http://www.omdbapi.com/?i=tt3896198&apikey=298f93b0"

// Get movie data from api

func main() {
	// todo: welcome message
	fmt.Println("Hi! Search for a Movie!")
	// todo: enter movie name to search for
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input a movie to search for")
	// scanner.Scan()
	input := scanner.Text()
	fmt.Printf("input text: %q\n", input)

	// todo: display search result
	res, err := http.Get(testUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	fmt.Printf("response: %d\n", res.StatusCode)

	var resStr strings.Builder

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	byteCount, _ := resStr.Write(body)

	fmt.Printf("byte count %d\n", byteCount)

	fmt.Printf("resStr %s\n", resStr.String())

	// todo: if no results display message and prompt again

	// todo: if more than one page provide next page option

	// todo: provide goto page option

	// todo: provide new search option
}
