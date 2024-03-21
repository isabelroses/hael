package lib

func GenderTurnary(gender string, f string, m string, b string) string {
	var out string
	switch gender {
	case "fem":
		out = f
	case "masc":
		out = m
	default:
		out = b
	}
	return out
}
