package main

import (
	"context"
	"fmt"
	"llm-engineer/pkg/llm"
	"log"
)

func main() {

	client := llm.NewClient("http://localhost:11434")

	ctx := context.Background()

	prompt := "Explain the concept of 'Recursion' to a junior developer in one sentence."

	// Configuatrion A: High Determinism (Code)
	tempLow := 0.1

	reqLow := llm.GenerateRequest{
		Model: "llama3",
		Prompt: prompt,
		Stream: false,
		Options: &llm.Options{
			Temperature: &tempLow,
		},
	}
	
	// Configuration B: High Creativity (Brainstorming)
	tempHigh := 0.9

	reqHigh := llm.GenerateRequest{
		Model: "llama3",
		Prompt: prompt,
		Stream: false,
		Options: &llm.Options{
			Temperature: &tempHigh,
		},
	}

	fmt.Println("=== Generating Report ===")

	// Execute Run A
	resLow, err := client.Generate(ctx, reqLow)
	if err != nil {
		log.Fatalf("Error (Low Temp): %v", err)
	}

	// Execute Run B
	resHigh, err := client.Generate(ctx, reqHigh)
	if err != nil {
		log.Fatalf("Error (High Temp): %v", err)
	}

	// The Report
	fmt.Println("\n=== PROMPT COMPARISON REPORT ===")
	fmt.Printf("Prompt: %s\n\n", prompt)
	fmt.Printf("[Temp 0.1]: %s\n", resLow.Response)
	fmt.Printf("[Temp 0.9]: %s\n", resHigh.Response)
	fmt.Println("================================")
}