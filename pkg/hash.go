package root

type Hash interface {
	Generate(s string) (string, error)
	Compare(hash, s string) error
}
