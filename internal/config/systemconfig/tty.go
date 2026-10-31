package systemconfig

import (
	"os"

	"github.com/git-town/git-town/v24/internal/config/configdomain"
	"github.com/mattn/go-isatty"
)

// DetermineStdioIsTerminal reports whether both STDIN and STDOUT are connected to a terminal.
// Unlike DetermineTTY, this doesn't fall back to opening the terminal device,
// because subprocesses only see the STDIN and STDOUT handles that Git Town passes to them.
func DetermineStdioIsTerminal() configdomain.StdioIsTerminal {
	return configdomain.StdioIsTerminal(isTerminal(os.Stdin) && isTerminal(os.Stdout))
}

// DetermineTTY reports whether an interactive terminal is available.
func DetermineTTY() configdomain.HasTTY {
	if isTerminal(os.Stdin) {
		return true
	}
	return canOpenTTY()
}

// isTerminal indicates whether the given file is connected to a terminal.
func isTerminal(file *os.File) bool {
	fd := file.Fd()
	return isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd)
}
