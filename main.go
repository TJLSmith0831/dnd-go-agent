package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tjlsmith0831/dnd-go-agent/pkg/llm"
	"github.com/tjlsmith0831/dnd-go-agent/pkg/session"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, using environment variables")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan string)
	out := make(chan string)
	go session.Run(ctx, in, out, llm.NewAnthropicClient())

	fmt.Println("DM mode active. Describe your action (ctrl+c to quit):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		in <- line
		fmt.Printf("\nNarrator: %s\n\n", <-out)
	}
}
