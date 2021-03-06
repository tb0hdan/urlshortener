package miscellaneous

// BuildInfo - store build information (for redistributable binaries)
type BuildInfo struct {
	Build,
	BuildDate,
	GoVersion,
	Version string
}

// ServerConfig - store HTTP server configuration
type ServerConfig struct {
	BuildInfo         *BuildInfo
	Bind              string
	ReadTimeout       int
	WriteTimeout      int
	URLDBFile         string
	TemplateDirectory string
}
