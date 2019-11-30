package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os/user"
	"strconv"
	"time"
)

func greet(name string, message string) string {
	return "Hello, " + name + "!\n" + message + "\n"
}

func whoami() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	return user.Name
}

func randomNum(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}

func printEastQuoteOnRandom() {
	type Quote struct {
		Author string `json:"author"`
		Book   string `json:"book"`
		Quote  string `json:"quote"`
	}

	type Quotes struct {
		Quotes []Quote `json:"quotes"`
	}

	byteValue, err := ioutil.ReadFile("./data.json")
	var eastQuotes Quotes
	json.Unmarshal(byteValue, &eastQuotes)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("\"" + eastQuotes.Quotes[randomNum(len(eastQuotes.Quotes))].Quote + "\"")
}

func printWestQuote() {
	type Quote struct {
		Quote string `json:"quote"`
	}

	resp, err := http.Get("https://api.kanye.rest")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var westQuotes Quote
	json.Unmarshal(body, &westQuotes)
	fmt.Println("\"" + westQuotes.Quote + "\"")
}

func questionGetPoints(correct string) bool {
	var reply string
	fmt.Print("Kanye West or east quote?!? (west) for Kanye West; (east) for Eastern Quotes:")
	for {
		fmt.Scan(&reply)
		if reply == "west" || reply == "east" {
			break
		} else {
			fmt.Println("Please, select type only \"west\" or \"east\"...:")
		}
	}

	if reply == correct {
		return true
	} else {
		return false
	}
}

func emoji(score int) string {
	if score < 4 {
		return ":("
	} else if score > 6 {
		return ":)"
	} else {
		return ":/"
	}
}

func main() {
	message := "This simple but yet deadly test can determine how well can you distinguish between Kenye West quotes and famous quotes from the Eastern culture.\nYou will be tested by trying to determine if a quote has been said by the self proclaimed genius or actual geniuses from the east.\nEast or West, chose wisely ;)"

	greatings := greet(whoami(), message)
	fmt.Println(greatings)
	time.Sleep(3 * time.Second)
	fmt.Println("Are you ready to start?!?")
	time.Sleep(3 * time.Second)
	fmt.Println("3...calculating the meaning of life...")
	time.Sleep(3 * time.Second)
	fmt.Println("2...reversing entropy...")
	time.Sleep(3 * time.Second)
	fmt.Println("1...dividing by zero...\n\n")

	var score int
	for i := 1; i <= 10; i++ {
		fmt.Println("Question #" + strconv.Itoa(i))
		if randomNum(2) == 0 {
			printEastQuoteOnRandom()
			if questionGetPoints("east") {
				fmt.Println("...True! Quote by:")
				score++
			}
		} else {
			printWestQuote()
			if questionGetPoints("west") {
				fmt.Println("...True! The quote is by the TRUE genius mastermind of the generation...")
				score++
			}
		}
	}
	fmt.Println("Your score is...")
	time.Sleep(1 * time.Second)
	fmt.Println("3...")
	time.Sleep(1 * time.Second)
	fmt.Println("2...")
	time.Sleep(1 * time.Second)
	fmt.Println("1...")
	fmt.Println(strconv.Itoa(score) + " out of 10! " + emoji(score))
}
