# Dev-util

A simple CLI tool to manage and start development servers for your projects from anywhere. No more `cd`-ing into project directories and remembering different commands!

## Features

- 🚀 **Quick dev server startup** - Start any registered project's dev server from anywhere
- 📝 **Project management** - Add, list, and remove projects easily
- 💾 **Persistent storage** - Projects are saved locally and persist between sessions
- 🎯 **Simple commands** - Intuitive CLI interface
- 🔧 **Flexible** - Works with any dev server command (npm, yarn, go, python, etc.)
- ⚡ **Smart autocomplete** - Tab completion for commands and project names
- 🎨 **Shell integration** - Automatic completion setup for bash, zsh, and fish

## Installation

### Option 1: Easy Installation (Recommended)

**One-command installation with automatic completion setup:**

```bash
./install.sh
```

This will:
- ✅ Build the binary
- ✅ Install to your system (system-wide or user directory)
- ✅ Set up shell completion automatically
- ✅ Configure your shell (bash/zsh/fish)
- ✅ Handle PATH configuration

### Option 2: Manual Installation

1. **Clone or download this repository:**
   ```bash
   git clone <your-repo-url>
   cd dev-util
   ```

2. **Install dependencies:**
   ```bash
   make deps
   ```

3. **Build and install:**
   ```bash
   # User installation (no sudo required)
   make install-user
   
   # OR system-wide installation (requires sudo)
   make install
   ```

4. **Setup completion (if not done automatically):**
   ```bash
   make setup-completion
   ```

### Option 3: Manual Build

1. **Download dependencies:**
   ```bash
   go mod download
   ```

2. **Build:**
   ```bash
   go build -o dev .
   ```

3. **Install:**
   ```bash
   # Copy to a directory in your PATH
   sudo cp dev /usr/local/bin/
   # OR
   cp dev ~/.local/bin/
   ```

4. **Setup completion manually:**
   ```bash
   # Generate completion script
   dev completion bash > ~/.local/share/bash-completion/completions/dev
   # Add to your shell config file
   ```

## Usage

### Adding Projects

Register a new project with its path and dev server command:

```bash
# Basic usage
dev add zensight-fe /path/to/zensight-fe "npm run dev"

# With description
dev add api-server /home/user/api "go run main.go" --description "Backend API server"

# Different types of projects
dev add frontend ./frontend "yarn start"
dev add backend ../backend "python manage.py runserver"
dev add mobile ./mobile "expo start"
```

### Starting Dev Servers

Start any registered project's dev server:

```bash
# Start the dev server
dev run zensight-fe

# The tool will:
# 1. Change to the project directory
# 2. Run the configured command
# 3. Display output in real-time
```

### Managing Projects

List all registered projects:

```bash
dev list
```

Remove a project:

```bash
# With confirmation prompt
dev remove zensight-fe

# Force remove without confirmation
dev remove zensight-fe --force
```

### Getting Help

```bash
# Show all commands
dev --help

# Show help for specific command
dev add --help
dev run --help
```

### Autocomplete

The tool includes smart autocomplete for commands and project names:

```bash
# Tab completion for main commands
dev <TAB>          # Shows: add, list, remove, run
dev run <TAB>       # Shows your registered projects
dev remove <TAB>    # Shows your registered projects

# Type to filter results
dev run zen<TAB>    # Shows only projects starting with "zen"
```

**Features:**
- ⚡ **Smart filtering** - Type letters to narrow down project names
- 🔄 **Cycling completion** - Each tab press shows the next option
- 🎯 **Context-aware** - Different completions for different commands
- 🚀 **Fast** - Completions load instantly from your project list

## Examples

### Setting up a typical workflow

1. **Add your projects:**
   ```bash
   dev add zensight-fe /home/user/projects/zensight-fe "npm run dev"
   dev add zensight-api /home/user/projects/zensight-api "go run main.go"
   dev add mobile-app /home/user/projects/mobile "expo start"
   ```

2. **List your projects:**
   ```bash
   dev list
   ```

3. **Start any project:**
   ```bash
   dev run zensight-fe
   ```

### Working with different project types

```bash
# React/Next.js projects
dev add my-react-app ./react-app "npm start"
dev add my-next-app ./next-app "npm run dev"

# Node.js/Express APIs
dev add my-api ./api "node server.js"
dev add my-express ./express-app "npm run dev"

# Go projects
dev add my-go-api ./go-api "go run main.go"
dev add my-go-service ./go-service "go run cmd/server/main.go"

# Python projects
dev add my-django ./django-app "python manage.py runserver"
dev add my-flask ./flask-app "flask run"

# Mobile development
dev add my-react-native ./rn-app "npx react-native start"
dev add my-expo ./expo-app "expo start"
```

## Configuration

Projects are stored in `~/.dev-util/projects.json`. This file contains all your registered projects and their configurations.

The configuration file is automatically created when you add your first project.

## Troubleshooting

### Command not found

If you get "command not found" after installation:

1. **Check your PATH:**
   ```bash
   echo $PATH
   ```

2. **For user installation, ensure ~/.local/bin is in PATH:**
   ```bash
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
   source ~/.bashrc
   ```

3. **For system installation, ensure /usr/local/bin is in PATH:**
   ```bash
   echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.bashrc
   source ~/.bashrc
   ```

### Project directory not found

If you get "Directory does not exist" errors:

1. **Check the path is correct:**
   ```bash
   ls -la /path/to/your/project
   ```

2. **Use absolute paths when adding projects:**
   ```bash
   dev add my-project /home/user/absolute/path/to/project "npm run dev"
   ```

### Permission issues

If you get permission errors:

1. **For system installation, use sudo:**
   ```bash
   sudo make install
   ```

2. **For user installation, ensure ~/.local/bin exists:**
   ```bash
   mkdir -p ~/.local/bin
   ```

## Development

### Building from source

```bash
# Download dependencies
make deps

# Build
make build

# Build for all platforms
make build-all

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean
```

### Project Structure

```
dev-util/
├── cmd/           # CLI commands
│   ├── root.go    # Root command
│   ├── add.go     # Add project command
│   ├── run.go     # Run project command
│   ├── list.go    # List projects command
│   └── remove.go  # Remove project command
├── models/        # Data models
│   └── project.go # Project model
├── storage/       # Data persistence
│   └── storage.go # Storage operations
├── main.go        # Application entry point
├── go.mod         # Go module file
├── Makefile       # Build automation
└── README.md      # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Support

If you encounter any issues or have questions:

1. Check the troubleshooting section above
2. Open an issue on GitHub
3. Check the help command: `dev --help`
