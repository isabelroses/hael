package lib

var (
	// TODO: expand on these arrays
	Fem = [12]string{
		"good girl",
		"you are a beacon of femininity, strength, and grace",
		"you exude elegance effortlessly",
		"you are a queen",
		"you are a goddess",
		"slayyy queen",
		"in your authenticity, you are a symphony of courage, resilience, and boundless beauty",
		"be yourself and your radiance will inspire others to do the same",
		"you are not just a star; you are the entire galaxy",
		"shine bright, darling; your authenticity is a sparkle in the universe",
		"embrace your hues, precious gem; you're a kaleidoscope of beauty",
		"blossom, darling bud; your petals whisper tales of resilience",
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
