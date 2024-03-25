package lib

var (
	// TODO: expand on these arrays
	Fem = [16]string{
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
		"you're an amazing girl <3",
		"you're beautiful",
		"you're loved",
		"You look beautiful just the way you are",
	}

	Masc = [14]string{
		"good boy",
		"you are a beacon of masculinity, strength, and power",
		"you exude ruggedness effortlessly",
		"you are a king",
		"rise and conquer, noble warrior; your strength knows no bounds",
		"you are the epitome of resilience, courage, and honor, standing tall amidst life's challenges",
		"in your presence, there's an allure of raw masculinity and unyielding determination",
		"you are not just a star; you're a blazing comet, carving through the cosmos with purpose and vigor",
		"with each step, you leave imprints of resilience and fortitude upon the earth",
		"stand tall like the mighty oak, weathering storms with unwavering strength and dignity",
		"you're an amazing boy <3",
		"you're beautiful",
		"you're loved",
		"You look handsome just the way you are",
	}

	Both = [(len(Fem) + len(Masc))]string{}
)

// setup sets up the both array with the values of the fem and masc arrays
func Setup() {
	copy(Both[:], Fem[:])
	copy(Both[len(Fem):], Masc[:])
}
