# LLM Engineer

A Go-based CLI application for exploring and comparing Large Language Model (LLM) responses using Ollama. This project demonstrates how different temperature settings affect model output, showcasing the balance between deterministic and creative responses.

## Overview

This project provides a simple yet powerful interface to interact with local LLM models through Ollama. It compares responses generated with different temperature settings:
- **Low Temperature (0.1)**: Deterministic, focused responses ideal for code generation and factual answers
- **High Temperature (0.9)**: Creative, varied responses perfect for brainstorming and diverse perspectives

## Features

- 🚀 Simple HTTP client for Ollama API
- 🎯 Temperature-based response comparison
- ⚙️ Configurable model parameters
- 🔄 Context-aware request handling
- 📊 Side-by-side output comparison

## Prerequisites

- **Go**: Version 1.24.3 or higher ([Download Go](https://go.dev/dl/))
- **Ollama**: Local LLM runtime ([Install Ollama](https://ollama.ai))
- **Homebrew** (for macOS installation): [Install Homebrew](https://brew.sh)

## Installation

### 1. Clone the Repository

```bash
git clone <your-repository-url>
cd llm
```

### 2. Install Ollama and Setup Model

The project includes a Makefile for easy setup:

```bash
# Install Ollama (macOS)
make install-ollama

# Start Ollama service
make start-ollama

# Pull the Llama 3 model
make pull-llama3
```

**Manual Installation:**
```bash
# Install Ollama
brew install ollama

# Start Ollama
ollama serve

# Pull Llama 3 model (in another terminal)
ollama pull llama3
```

### 3. Install Go Dependencies

```bash
go mod download
```

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── pkg/
│   └── llm/
│       ├── client.go     # HTTP client implementation
│       └── types.go      # Request/Response types
├── go.mod                # Go module definition
├── Makefile              # Build and setup automation
├── README.md             # This file
└── .gitignore            # Git ignore rules
```

## Usage

### Running the Application

```bash
# Run directly with Go
go run cmd/main.go

# Or build and execute
go build -o bin/llm cmd/main.go
./bin/llm
```

### Expected Output

```
=== Generating Report ===

=== PROMPT COMPARISON REPORT ===
Prompt: Explain the concept of 'Recursion' to a junior developer in one sentence.

[Temp 0.1]: Recursion is when a function calls itself repeatedly until...
[Temp 0.9]: Imagine a function that's like a mirror reflecting itself, creating...
================================
```

## Configuration

### Modifying the Prompt

Edit the `prompt` variable in [cmd/main.go](cmd/main.go):

```go
prompt := "Your custom prompt here"
```

### Adjusting Temperature Settings

Temperature controls randomness (0.0 = deterministic, 1.0 = creative):

```go
tempLow := 0.1   // More focused
tempHigh := 0.9  // More creative
```

### Changing the Model

Modify the `Model` field in the request:

```go
reqLow := llm.GenerateRequest{
    Model: "llama3",  // Change to "codellama", "mistral", etc.
    // ...
}
```

### Available Models

Check available models:
```bash
ollama list
```

Pull additional models:
```bash
ollama pull codellama
ollama pull mistral
```

## API Reference

### Client Creation

```go
client := llm.NewClient("http://localhost:11434")
```

### Generate Request

```go
req := llm.GenerateRequest{
    Model:   "llama3",
    Prompt:  "Your prompt",
    Stream:  false,
    Options: &llm.Options{
        Temperature: &temp,
    },
}

response, err := client.Generate(ctx, req)
```

## Troubleshooting

### Ollama Not Running

```bash
# Check if Ollama is running
curl http://localhost:11434/api/version

# Start Ollama
make start-ollama
```

### Model Not Found

```bash
# List installed models
ollama list

# Pull required model
ollama pull llama3
```

### Request Timeout

Increase the timeout in [pkg/llm/client.go](pkg/llm/client.go):

```go
HttpClient: &http.Client{
    Timeout: 120 * time.Second, // Increase from 60s
}
```

## Development

### Building

```bash
go build -o bin/llm cmd/main.go
```

### Testing

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.

## Acknowledgments

- [Ollama](https://ollama.ai) for providing local LLM runtime
- [Llama 3](https://ai.meta.com/llama/) by Meta AI
- Go community for excellent tooling and libraries

## Contact

For questions or feedback, please open an issue on GitHub.

---

**Happy Coding! 🚀**