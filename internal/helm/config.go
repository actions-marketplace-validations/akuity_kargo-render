package helm

// Config encapsulates optional Helm configuration options.
type Config struct {
	// ReleaseName specifies the release name that will be used when executing the
	// `helm template` command.
	ReleaseName string `json:"releaseName,omitempty"`
	// ChartPath is a path to a directory, relative to the root of the repository,
	// where a Helm chart can be located. This is used as an argument in the
	// `helm template` command. By convention, if left unspecified, the value
	// `base/` is assumed.
	ChartPath string `json:"chartPath,omitempty"`
	// Values are paths to Helm values files (e.g. values.yaml), relative to the
	// root of the repository. Each of these will be used as a value for the
	// `--values` flag in the `helm template` command. By convention, if left
	// unspecified, one path will be assumed: <branch name>/values.yaml.
	ValuesPaths []string `json:"valuesPaths,omitempty"`
}