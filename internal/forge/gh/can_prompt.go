package gh

import (
	"github.com/git-town/git-town/v24/internal/config/configdomain"
	"github.com/git-town/git-town/v24/internal/forge/forgedomain"
)

// CanPrompt indicates whether gh may prompt the user for missing data.
// This is the case if Git Town runs interactively and gh is able to prompt.
// The latter mirrors IOStreams.CanPrompt in gh: STDIN and STDOUT must be terminals
// and GH_PROMPT_DISABLED must not be set.
// It ignores the "prompt" setting in the gh config.
type CanPrompt bool

// NewCanPrompt determines whether gh may prompt the user, given the current environment.
func NewCanPrompt(interactive configdomain.Interactive, stdioIsTerminal configdomain.StdioIsTerminal, promptDisabled forgedomain.GhPromptDisabled) CanPrompt {
	return CanPrompt(interactive.IsEnabled() && bool(stdioIsTerminal) && !bool(promptDisabled))
}
