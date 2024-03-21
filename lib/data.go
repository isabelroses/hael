package lib

var (
	// TODO: expand on these arrays
	Fem = [6]string{
		"good girl",
		"you are a beacon of femininity, strength, and grace",
		"you exude elegance effortlessly",
		"you are a queen",
		"you are a goddess",
		"slayyy queen",
	}

	Masc = [4]string{
		"good boy",
		"you are a beacon of masculinity, strength, and power",
		"you exude ruggedness effortlessly",
		"you are a king",
	}

	Both = [(len(Fem) + len(Masc))]string{}
)

// setup sets up the both array with the values of the fem and masc arrays
func Setup() {
	copy(Both[:], Fem[:])
	copy(Both[len(Fem):], Masc[:])
}
