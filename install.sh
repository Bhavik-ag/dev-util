#!/bin/bash

# Dev-util Installation Script

set -e

echo "🚀 Installing dev-util..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed. Please install Go first.${NC}"
    echo "Visit: https://golang.org/doc/install"
    exit 1
fi

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "📦 Downloading dependencies..."
go mod tidy

echo "🔨 Building dev-util..."
go build -o dev .

echo "📁 Creating build directory..."
mkdir -p build
mv dev build/

echo "🔧 Installing dev-util..."

# Check if user wants system-wide or user installation
echo "Choose installation method:"
echo "1) System-wide installation (requires sudo, installs to /usr/local/bin)"
echo "2) User installation (no sudo required, installs to ~/.local/bin)"
echo "3) Skip installation (just build the binary)"

read -p "Enter your choice (1-3): " choice

case $choice in
    1)
        echo "Installing system-wide..."
        sudo cp build/dev /usr/local/bin/
        echo -e "${GREEN}✅ dev-util installed system-wide to /usr/local/bin${NC}"
        ;;
    2)
        echo "Installing to user directory..."
        mkdir -p ~/.local/bin
        cp build/dev ~/.local/bin/
        
        # Check if ~/.local/bin is in PATH
        if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
            echo -e "${YELLOW}⚠️  ~/.local/bin is not in your PATH. Adding it to ~/.bashrc...${NC}"
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
            echo -e "${YELLOW}Please run 'source ~/.bashrc' or restart your terminal.${NC}"
        fi
        
        echo -e "${GREEN}✅ dev-util installed to ~/.local/bin${NC}"
        ;;
    3)
        echo -e "${GREEN}✅ Binary built successfully in build/dev${NC}"
        echo "You can manually install it later or run it directly:"
        echo "  ./build/dev --help"
        ;;
    *)
        echo -e "${RED}Invalid choice. Installation cancelled.${NC}"
        exit 1
        ;;
esac

echo ""
echo -e "${GREEN}🎉 Installation complete!${NC}"

# Setup shell completion
echo ""
echo -e "${BLUE}🔧 Setting up shell completion...${NC}"

# Detect shell
SHELL_NAME=$(basename "$SHELL")
COMPLETION_DIR=""

case "$SHELL_NAME" in
    "bash")
        COMPLETION_DIR="$HOME/.local/share/bash-completion/completions"
        COMPLETION_FILE="$COMPLETION_DIR/dev"
        ;;
    "zsh")
        COMPLETION_DIR="$HOME/.local/share/zsh/site-functions"
        COMPLETION_FILE="$COMPLETION_DIR/_dev"
        ;;
    "fish")
        COMPLETION_DIR="$HOME/.config/fish/completions"
        COMPLETION_FILE="$COMPLETION_DIR/dev.fish"
        ;;
    *)
        echo -e "${YELLOW}⚠️  Unsupported shell '$SHELL_NAME'. Skipping completion setup.${NC}"
        echo -e "${YELLOW}Supported shells: bash, zsh, fish${NC}"
        COMPLETION_DIR=""
        ;;
esac

if [ -n "$COMPLETION_DIR" ]; then
    echo -e "${YELLOW}Detected shell: $SHELL_NAME${NC}"
    
    # Create completion directory
    mkdir -p "$COMPLETION_DIR"
    
    # Generate completion script
    echo -e "${YELLOW}Generating completion script...${NC}"
    
    case "$SHELL_NAME" in
        "bash")
            ./build/dev completion bash > "$COMPLETION_FILE"
            echo -e "${GREEN}✅ Bash completion installed to: $COMPLETION_FILE${NC}"
            
            # Add completion setup to .bashrc if not already there
            if ! grep -q "dev-util" ~/.bashrc 2>/dev/null; then
                echo "" >> ~/.bashrc
                echo "# Enable bash completion for dev-util" >> ~/.bashrc
                echo "if [ -d ~/.local/share/bash-completion/completions ]; then" >> ~/.bashrc
                echo "    for file in ~/.local/share/bash-completion/completions/*; do" >> ~/.bashrc
                echo "        [ -r \"\$file\" ] && source \"\$file\"" >> ~/.bashrc
                echo "    done" >> ~/.bashrc
                echo "fi" >> ~/.bashrc
                echo -e "${YELLOW}⚠️  Added bash completion setup to ~/.bashrc${NC}"
            else
                echo -e "${GREEN}✅ Bash completion already configured in ~/.bashrc${NC}"
            fi
            
            # Add shell integration (dev-cd function) to .bashrc
            if ! grep -q 'eval "$(dev init bash)"' ~/.bashrc 2>/dev/null; then
                echo "" >> ~/.bashrc
                echo "# dev-util shell integration (enables dev-cd command)" >> ~/.bashrc
                echo 'eval "$(dev init bash)"' >> ~/.bashrc
                echo -e "${GREEN}✅ Shell integration added to ~/.bashrc${NC}"
            else
                echo -e "${GREEN}✅ Shell integration already configured in ~/.bashrc${NC}"
            fi
            ;;
        "zsh")
            ./build/dev completion zsh > "$COMPLETION_FILE"
            echo -e "${GREEN}✅ Zsh completion installed to: $COMPLETION_FILE${NC}"
            
            # Add completion setup to .zshrc if not already there
            if ! grep -q "fpath.*site-functions" ~/.zshrc 2>/dev/null; then
                echo "" >> ~/.zshrc
                echo "# Enable zsh completion" >> ~/.zshrc
                echo "fpath=(~/.local/share/zsh/site-functions \$fpath)" >> ~/.zshrc
                echo "autoload -U compinit && compinit" >> ~/.zshrc
                echo -e "${YELLOW}⚠️  Added zsh completion setup to ~/.zshrc${NC}"
            fi
            
            # Add shell integration (dev-cd function) to .zshrc
            if ! grep -q 'eval "$(dev init zsh)"' ~/.zshrc 2>/dev/null; then
                echo "" >> ~/.zshrc
                echo "# dev-util shell integration (enables dev-cd command)" >> ~/.zshrc
                echo 'eval "$(dev init zsh)"' >> ~/.zshrc
                echo -e "${GREEN}✅ Shell integration added to ~/.zshrc${NC}"
            else
                echo -e "${GREEN}✅ Shell integration already configured in ~/.zshrc${NC}"
            fi
            ;;
        "fish")
            ./build/dev completion fish > "$COMPLETION_FILE"
            echo -e "${GREEN}✅ Fish completion installed to: $COMPLETION_FILE${NC}"
            
            # Add shell integration to fish config
            FISH_CONFIG="$HOME/.config/fish/config.fish"
            mkdir -p "$(dirname "$FISH_CONFIG")"
            if ! grep -q "dev init fish" "$FISH_CONFIG" 2>/dev/null; then
                echo "" >> "$FISH_CONFIG"
                echo "# dev-util shell integration (enables dev-cd command)" >> "$FISH_CONFIG"
                echo "dev init fish | source" >> "$FISH_CONFIG"
                echo -e "${GREEN}✅ Shell integration added to $FISH_CONFIG${NC}"
            else
                echo -e "${GREEN}✅ Shell integration already configured in $FISH_CONFIG${NC}"
            fi
            ;;
    esac
    
    echo ""
    echo -e "${GREEN}🎉 Shell completion and integration setup complete!${NC}"
    echo ""
    echo -e "${YELLOW}To activate:${NC}"
    case "$SHELL_NAME" in
        "bash")
            echo "  source ~/.bashrc"
            echo "  # OR restart your terminal"
            ;;
        "zsh")
            echo "  source ~/.zshrc"
            echo "  # OR restart your terminal"
            ;;
        "fish")
            echo "  # Restart your terminal or run: source ~/.config/fish/config.fish"
            ;;
    esac
    echo ""
    echo -e "${YELLOW}Available commands:${NC}"
    echo "  dev-cd <project>   - Change to a project directory"
    echo "  dev-run <project>  - Run a project's dev server"
    echo "  dev <TAB>          - Tab completion for all commands"
fi

echo ""
echo "Usage examples:"
echo "  dev add my-project /path/to/project 'npm run dev'"
echo "  dev list"
echo "  dev run my-project"
echo "  dev --help"
echo ""
echo "For more information, see the README.md file."
