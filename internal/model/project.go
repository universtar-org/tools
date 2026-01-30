package model

// Username can be get from the file name, so it is unnecessary in this structure.
type Project struct {
	Repo        string   `yaml:"repo"`
	Stars       int      `yaml:"stars"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	UpdatedAt   string   `yaml:"updated_at"`
}
