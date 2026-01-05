.PHONY: install-ollama start-ollama pull-llama3 help

help:
	@echo "Available targets:"
	@echo "  install-ollama  - Install Ollama on macOS"
	@echo "  start-ollama    - Start the Ollama service"
	@echo "  pull-llama3     - Pull the Llama 3 model using Ollama"

install-ollama:
	@echo "Installing Ollama via Homebrew..."
	@which brew > /dev/null || (echo "Error: Homebrew is not installed. Install it from https://brew.sh" && exit 1)
	brew install ollama

start-ollama:
	@echo "Starting Ollama service..."
	@open -a Ollama || ollama serve > /dev/null 2>&1 &
	@echo "Waiting for Ollama to start..."
	@sleep 3

pull-llama3: start-ollama
	@echo "Pulling Llama 3 model..."
	ollama pull llama3
