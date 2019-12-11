package miscellaneous

type BuildInfo struct {
	Build,
	BuildDate,
	GoVersion,
	Version string
}

type ServerConfig struct {
	BuildInfo         *BuildInfo
	Bind              string
	ReadTimeout       int
	WriteTimeout      int
	URLDBFile         string
	TemplateDirectory string
}
