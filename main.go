package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	ctx := context.Background()
	client, err := genai.NewClient(
		ctx,
		&genai.ClientConfig{
			APIKey:  os.Getenv("GEMINI_API_KEY"),
			Backend: genai.BackendGeminiAPI,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("gemini>\nyou> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" || input == "quit" {
			break
		}
		result, err := client.Models.GenerateContent(
			ctx,
			"gemini-2.0-flash",
			genai.Text(input),
			nil,
		)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result.Text())
		}
		fmt.Print("you> ")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
