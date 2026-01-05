package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	// Import your local packages
	"llm-engineer/pkg/llm"
	"llm-engineer/pkg/resume"
)

// constructPrompt builds the string we send to the LLM.
// Senior Tip: Keep prompt logic separate from business logic.
func constructPrompt(jobDescription string) string {
    // TODO: Write a prompt that asks the LLM to extract key skills
    // and generate 3 professional resume bullet points.
	prompt := `You are an expert resume writer. Given the following job description, 
	extract the key skills and technologies required for the role. 
	Then, generate 3 concise and impactful resume bullet points that highlight relevant experience and achievements using those skills.
	Please respond in the following JSON format:

	{
	"bullets": [
		"First bullet point here.",
		"Second bullet point here.",
		"Third bullet point here."
	]
	}

	Job Description:
	%s
	Remember to focus on clarity, relevance, and quantifiable achievements in the bullet points.`
    
	// CRITICAL: You MUST tell the LLM exactly what JSON schema to use 
    // in the text prompt as well, so it knows the field names.
    return fmt.Sprintf(prompt, jobDescription)
}

func main() {
	// 1. Setup
	client := llm.NewClient("http://localhost:11434")
	ctx := context.Background()

	// 2. Input Data (Simulated for now)
	rawJobDescription := "We are looking for a Software Engineer with Go experience. Must know Docker, Kubernetes, and how to build scalable microservices."
	temp := 0.9
	// 3. Configuration
    // TODO: Create the GenerateRequest.
	propmt := constructPrompt(rawJobDescription)
	req := llm.GenerateRequest{
		Model: "llama3",
		Prompt: propmt,
		Stream: false,
		Options: &llm.Options{
			Temperature: &temp,
		},
		Format: "json",
	}
    // IMPORTANT: specificy Format: "json" here!
    // Set stream to false for simplicity.
	
	fmt.Println("Generating resume bullets...")

	// 4. Execution
	// TODO: Call client.Generate()
	res, err := client.Generate(ctx, req)

	

	// 5. Handling Response
    // TODO: Check for errors.
	if err != nil {
		log.Fatalf("Error:  %v", err)
	}
	
    // 6. Unmarshaling (The Moment of Truth)
	var result resume.GeneratedBullets

    
    // TODO: Use json.Unmarshal([]byte(resp.Response), &result)
    // Note: resp.Response is a string, Unmarshal wants bytes: []byte(string)

	json.Unmarshal( []byte(res.Response), &result)
	
	// 7. Output
	fmt.Printf("Successfully generated %d bullets:\n", len(result.Bullets))
	for _, b := range result.Bullets {
		fmt.Printf("- %s\n", b)
	}
}