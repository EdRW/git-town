package gh

import (
	"os"

	"github.com/git-town/git-town/v24/internal/config/configdomain"
	"github.com/mattn/go-isatty"
)

// CanPrompt indicates whether gh may prompt the user for missing data.
// This is the case if Git Town runs interactively and gh is able to prompt.
// The latter mirrors IOStreams.CanPrompt in gh: STDIN and STDOUT must be terminals
// and GH_PROMPT_DISABLED must not be set (any value disables prompts).
// It ignores the "prompt" setting in the gh config.
type CanPrompt bool

// DetectCanPrompt determines whether gh may prompt the user in the current environment.
func DetectCanPrompt(interactive configdomain.Interactive) CanPrompt {
	if !interactive.IsEnabled() {
		return false
	}
	if _, disabled := os.LookupEnv("GH_PROMPT_DISABLED"); disabled {
		return false
	}
	return CanPrompt(isTerminal(os.Stdin) && isTerminal(os.Stdout))
}

func isTerminal(file *os.File) bool {
	fd := file.Fd()
	return isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd)
}
