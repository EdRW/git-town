package gh_test

import (
	"testing"

	"github.com/git-town/git-town/v24/internal/config/configdomain"
	"github.com/git-town/git-town/v24/internal/forge/forgedomain"
	"github.com/git-town/git-town/v24/internal/forge/gh"
	"github.com/shoenig/test/must"
)

func TestNewCanPrompt(t *testing.T) {
	t.Parallel()
	interactive := configdomain.InteractiveEnabled
	nonInteractive := configdomain.Interactive("disabled via CLI")
	tests := []struct {
		name            string
		interactive     configdomain.Interactive
		stdioIsTerminal configdomain.StdioIsTerminal
		promptDisabled  forgedomain.GhPromptDisabled
		want            gh.CanPrompt
	}{
		{name: "interactive, terminal, prompts enabled", interactive: interactive, stdioIsTerminal: true, promptDisabled: false, want: true},
		{name: "interactive, terminal, prompts disabled", interactive: interactive, stdioIsTerminal: true, promptDisabled: true, want: false},
		{name: "interactive, no terminal, prompts enabled", interactive: interactive, stdioIsTerminal: false, promptDisabled: false, want: false},
		{name: "interactive, no terminal, prompts disabled", interactive: interactive, stdioIsTerminal: false, promptDisabled: true, want: false},
		{name: "non-interactive, terminal, prompts enabled", interactive: nonInteractive, stdioIsTerminal: true, promptDisabled: false, want: false},
		{name: "non-interactive, terminal, prompts disabled", interactive: nonInteractive, stdioIsTerminal: true, promptDisabled: true, want: false},
		{name: "non-interactive, no terminal, prompts enabled", interactive: nonInteractive, stdioIsTerminal: false, promptDisabled: false, want: false},
		{name: "non-interactive, no terminal, prompts disabled", interactive: nonInteractive, stdioIsTerminal: false, promptDisabled: true, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			have := gh.NewCanPrompt(test.interactive, test.stdioIsTerminal, test.promptDisabled)
			must.EqOp(t, test.want, have)
		})
	}
}
