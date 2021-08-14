package parsers

type Parser interface {
	Parse(string) error
}
