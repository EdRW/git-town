package envconfig

import "github.com/git-town/git-town/v24/internal/forge/forgedomain"

// GhPromptDisabled indicates whether the given environment disables the prompts of the gh CLI.
// Like gh, any value of GH_PROMPT_DISABLED disables prompts, even an empty one.
func GhPromptDisabled(env EnvVars) forgedomain.GhPromptDisabled {
	return forgedomain.GhPromptDisabled(env.GetOpt(ghPromptDisabled).IsSome())
}
