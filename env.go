package cliflags

type Namer interface {
	Name(string) string
	Env(string) []string
}
