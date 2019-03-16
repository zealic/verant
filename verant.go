package verant

type Type int

const (
	// plain text
	TypeText SpecType = iota
	// Dockerfile
	TypeDockerfile SpecType = iota
	// shell script file
	TypeShell SpecType = iota
	// docker-compose file
	TypeDockerCompose SpecType = iota
)

type Verant struct {
	File  string
	Type  Type
	Specs []*VersionSpec
}

type VersionSpec struct {
	Name       string
	Version    string
	SourceFile string
	SourceLine int
	Raw        string
}

func New(file string) *Verant {
	panic("Not impl")
}

func ScanDir(dir string) []*Verant {
	panic("Not impl")
}
