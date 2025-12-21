package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const bashInit = `
# dev-util shell integration
# Add this to your ~/.bashrc or ~/.zshrc:
#   eval "$(dev init bash)"

# Change directory to a registered project
function dev-cd() {
    local result
    result="$(command dev cd --path "$@" 2>&1)"
    if [[ $? -eq 0 && -n "$result" ]]; then
        builtin cd -- "$result" || return 1
        echo "📁 Changed to: $result"
    else
        echo "$result" >&2
        return 1
    fi
}

# Run dev server for a project (optionally change to its directory first)
function dev-run() {
    command dev run "$@"
}

# Completions for dev-cd
if [[ -n "${BASH_VERSION:-}" ]]; then
    _dev_cd_completions() {
        local cur="${COMP_WORDS[COMP_CWORD]}"
        local projects
        projects="$(command dev list --names-only 2>/dev/null)"
        COMPREPLY=($(compgen -W "$projects" -- "$cur"))
    }
    complete -F _dev_cd_completions dev-cd
fi
`

const zshInit = `
# dev-util shell integration
# Add this to your ~/.zshrc:
#   eval "$(dev init zsh)"

# Change directory to a registered project
function dev-cd() {
    local result
    result="$(command dev cd --path "$@" 2>&1)"
    if [[ $? -eq 0 && -n "$result" ]]; then
        builtin cd -- "$result" || return 1
        echo "📁 Changed to: $result"
    else
        echo "$result" >&2
        return 1
    fi
}

# Run dev server for a project
function dev-run() {
    command dev run "$@"
}

# Completions for dev-cd
if (( $+commands[dev] )); then
    _dev_cd() {
        local projects
        projects=("${(@f)$(command dev list --names-only 2>/dev/null)}")
        _describe 'project' projects
    }
    compdef _dev_cd dev-cd
fi
`

const fishInit = `
# dev-util shell integration
# Add this to your ~/.config/fish/config.fish:
#   dev init fish | source

# Change directory to a registered project
function dev-cd
    set -l result (command dev cd --path $argv 2>&1)
    if test $status -eq 0 -a -n "$result"
        builtin cd -- $result
        and echo "📁 Changed to: $result"
    else
        echo $result >&2
        return 1
    end
end

# Run dev server for a project
function dev-run
    command dev run $argv
end

# Completions for dev-cd
complete -c dev-cd -f -a "(command dev list --names-only 2>/dev/null)"
`

var initCmd = &cobra.Command{
	Use:   "init [shell]",
	Short: "Generate shell integration code",
	Long: `Generate shell integration code for dev-util.
This enables the 'dev-cd' function to change directories without spawning a new shell.

Supported shells: bash, zsh, fish

Usage:
  # For bash, add to ~/.bashrc:
  eval "$(dev init bash)"

  # For zsh, add to ~/.zshrc:
  eval "$(dev init zsh)"

  # For fish, add to ~/.config/fish/config.fish:
  dev init fish | source

After setup, use 'dev-cd <project>' to navigate to a project directory.`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh", "fish"},
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]
		switch shell {
		case "bash":
			fmt.Print(bashInit)
		case "zsh":
			fmt.Print(zshInit)
		case "fish":
			fmt.Print(fishInit)
		default:
			fmt.Printf("Error: Unsupported shell '%s'. Supported shells: bash, zsh, fish\n", shell)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
