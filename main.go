package precommit

type Repo struct {
	Repo  string
	Rev   string `yaml:"rev,omitempty"`
	Hooks []Hook
}

type Hook struct {
	ID                      string
	Exclude                 string   `yaml:"exclude,omitempty"`
	Name                    string   `yaml:"name,omitempty"`
	MinimumPreCommitVersion string   `yaml:"minimum_pre_commit_version,omitempty"`
	Language                string   `yaml:"language,omitempty"`
	AdditionalDependencies  []string `yaml:"additional_dependencies,omitempty"`
	Entry                   string   `yaml:"entry,omitempty"`
	Args                    []string `yaml:"args,omitempty"`
	Types                   []string `yaml:"types,omitempty"`
}
