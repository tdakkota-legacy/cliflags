package cliflags

type Namer interface {
	Name(string) string
	Env(string) []string
}

type DefaultNamer struct{}

func (d DefaultNamer) Name(n string) string {
	return n
}

func (d DefaultNamer) Env(n string) []string {
	return []string{n}
}
