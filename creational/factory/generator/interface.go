package generator

type UrlGenerator interface {
	Generate(original string) string
}
