# Dev-util Makefile

BINARY_NAME=dev
BUILD_DIR=build
VERSION?=1.0.0

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	@echo "✅ Cross-platform builds complete"

# Install globally (requires sudo on Linux/macOS)
install: build
	@echo "Installing $(BINARY_NAME) globally..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "✅ $(BINARY_NAME) installed globally"
	@$(MAKE) setup-completion

# Install to user's local bin (no sudo required)
install-user: build
	@echo "Installing $(BINARY_NAME) to user bin..."
	@mkdir -p ~/.local/bin
	@cp $(BUILD_DIR)/$(BINARY_NAME) ~/.local/bin/
	@echo "✅ $(BINARY_NAME) installed to ~/.local/bin"
	@echo "Make sure ~/.local/bin is in your PATH"
	@$(MAKE) setup-completion

# Uninstall
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@rm -f ~/.local/bin/$(BINARY_NAME)
	@echo "✅ $(BINARY_NAME) uninstalled"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@golangci-lint run

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Setup shell completion
setup-completion:
	@echo "Setting up shell completion..."
	@if command -v dev >/dev/null 2>&1; then \
		SHELL_NAME=$$(basename "$$SHELL"); \
		case "$$SHELL_NAME" in \
			"bash") \
				mkdir -p ~/.local/share/bash-completion/completions; \
				dev completion bash > ~/.local/share/bash-completion/completions/dev; \
				if ! grep -q "dev-util" ~/.bashrc 2>/dev/null; then \
					echo "" >> ~/.bashrc; \
					echo "# Enable bash completion for dev-util" >> ~/.bashrc; \
					echo "if [ -d ~/.local/share/bash-completion/completions ]; then" >> ~/.bashrc; \
					echo "    for file in ~/.local/share/bash-completion/completions/*; do" >> ~/.bashrc; \
					echo "        [ -r \"\$$file\" ] && source \"\$$file\"" >> ~/.bashrc; \
					echo "    done" >> ~/.bashrc; \
					echo "fi" >> ~/.bashrc; \
					echo "✅ Bash completion configured"; \
				else \
					echo "✅ Bash completion already configured"; \
				fi; \
				;; \
			"zsh") \
				mkdir -p ~/.local/share/zsh/site-functions; \
				dev completion zsh > ~/.local/share/zsh/site-functions/_dev; \
				if ! grep -q "dev-util" ~/.zshrc 2>/dev/null; then \
					echo "" >> ~/.zshrc; \
					echo "# Enable zsh completion for dev-util" >> ~/.zshrc; \
					echo "fpath=(~/.local/share/zsh/site-functions \$$fpath)" >> ~/.zshrc; \
					echo "autoload -U compinit && compinit" >> ~/.zshrc; \
					echo "✅ Zsh completion configured"; \
				else \
					echo "✅ Zsh completion already configured"; \
				fi; \
				;; \
			"fish") \
				mkdir -p ~/.config/fish/completions; \
				dev completion fish > ~/.config/fish/completions/dev.fish; \
				echo "✅ Fish completion configured"; \
				;; \
			*) \
				echo "⚠️  Unsupported shell: $$SHELL_NAME. Supported: bash, zsh, fish"; \
				;; \
		esac; \
	else \
		echo "❌ dev command not found. Please install first."; \
	fi

# Show help
help:
	@echo "Available commands:"
	@echo "  build        - Build the binary"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  install      - Install globally (requires sudo)"
	@echo "  install-user - Install to user bin (~/.local/bin)"
	@echo "  uninstall    - Remove from system"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  deps         - Download dependencies"
	@echo "  setup-completion - Setup shell completion"
	@echo "  help         - Show this help"

.PHONY: build build-all install install-user uninstall clean test fmt lint deps setup-completion help
