package envconfig_test

import (
	"testing"

	"github.com/git-town/git-town/v24/internal/config/envconfig"
	"github.com/git-town/git-town/v24/internal/forge/forgedomain"
	"github.com/shoenig/test/must"
)

func TestGhPromptDisabled(t *testing.T) {
	t.Parallel()
	tests := map[string]forgedomain.GhPromptDisabled{
		"":                     false,
		"GH_PROMPT_DISABLED=":  true,
		"GH_PROMPT_DISABLED=0": true,
		"GH_PROMPT_DISABLED=1": true,
		"OTHER=1":              false,
	}
	for give, want := range tests {
		t.Run(give, func(t *testing.T) {
			t.Parallel()
			env := envconfig.NewEnvVars([]string{give})
			have := envconfig.GhPromptDisabled(env)
			must.EqOp(t, want, have)
		})
	}
}
