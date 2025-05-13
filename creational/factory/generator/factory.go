package generator

func Factory(method string) UrlGenerator {
	switch method {
	case "hash":
		return HashGenerator{}
	default:
		return RandomGenerator{}
	}
}