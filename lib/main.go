package lib

func GenderReduce(gender string, f string, m string, n string, d string) string {
	var out string
	switch gender {
	case "fem":
		out = f
	case "masc":
		out = m
	case "neutral":
		out = n
	default:
		out = d
	}
	return out
}

func GenderTernary(gender string, f string, m string, d string) string {
	var out string
	switch gender {
	case "fem":
		out = f
	case "masc":
		out = m
	default:
		out = d
	}
	return out
}
