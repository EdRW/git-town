package configdomain

// StdioIsTerminal indicates whether both STDIN and STDOUT of this process are connected to a terminal.
// Subprocesses inherit these handles, so this determines whether they can prompt the user.
type StdioIsTerminal bool
