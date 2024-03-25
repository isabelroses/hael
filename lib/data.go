package lib

var (
	// TODO: expand on these arrays
	Fem = [15]string{
		"good girl",
		"slayyy queen",
		"you are a queen",
		"you're beautiful",
		"you are a goddess",
		"you're an amazing girl <3",
		"you exude elegance effortlessly",
		"you look beautiful just the way you are",
		"you are not just a star; you are the entire galaxy",
		"you are a beacon of femininity, strength, and grace",
		"blossom, darling bud; your petals whisper tales of resilience",
		"embrace your hues, precious gem; you're a kaleidoscope of beauty",
		"be yourself and your radiance will inspire others to do the same",
		"shine bright, darling; your authenticity is a sparkle in the universe",
		"in your authenticity, you are a symphony of courage, resilience, and boundless beauty",
	}

	Masc = [13]string{
		"good boy",
		"you are a king",
		"you're handsome",
		"you're an amazing boy <3",
		"you exude ruggedness effortlessly",
		"you look handsome just the way you are",
		"you are a beacon of masculinity, strength, and power",
		"rise and conquer, noble warrior; your strength knows no bounds",
		"with each step, you leave imprints of resilience and fortitude upon the earth",
		"in your presence, there's an allure of raw masculinity and unyielding determination",
		"stand tall like the mighty oak, weathering storms with unwavering strength and dignity",
		"you are the epitome of resilience, courage, and honor, standing tall amidst life's challenges",
		"you are not just a star; you're a blazing comet, carving through the cosmos with purpose and vigor",
	}

	Neutral = [3]string{
		"you're loved",
		"you are valid",
		"you're amazing",
	}

	All = [(len(Fem) + len(Masc) + len(Neutral))]string{}
)

// setup sets up the both array with the values of the fem and masc arrays
func Setup() {
	copy(All[:], Fem[:])
	copy(All[len(Fem):], Masc[:])
	copy(All[len(Fem)+len(Masc):], Neutral[:])

	copy(Fem[:], Neutral[:])
	copy(Masc[:], Neutral[:])
}
