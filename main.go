package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/isabelroses/hael/lib"
)

var (
	// pronouns = os.Getenv("HAEL_PRONOUNS") TODO: something with pronouns
	gender = os.Getenv("HAEL_GENDER")

	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger

	fEuphoria bool
	fPickmeup bool
)

const helpMessage = `Usage: hael [FLAGS]

Flags:
  -h, --help    Print this help message
  -r, --random  Print a random euphoric message
  -p, --pickup  Print a pick-me-up message

Environment Variables:
  HAEL_GENDER  'fem', 'masc', 'neutral' or 'both'

Examples:
  hael -r
  hael --pickup

Help:
  https://github.com/isabelroses/hael`

func main() {
	lib.Setup()

	flag.BoolVar(&fEuphoria, "random", false, "Print a random euphoric message")
	flag.BoolVar(&fEuphoria, "r", false, "Print a random euphoric message")
	flag.BoolVar(&fPickmeup, "p", false, "Print a pick-me-up message")
	flag.BoolVar(&fPickmeup, "pickup", false, "Print a pick-me-up message")
	flag.Usage = func() {
		fmt.Println(helpMessage)
	}

	flag.Parse()

	switch {
	case fPickmeup:
		pickmeup()
	case fEuphoria:
		randomEuphoria()
	default:
		randomEuphoria()
	}
}

func randomEuphoria() {
	out := lib.GenderReduce(
		gender,
		lib.Fem[rand.Intn(len(lib.Fem))],
		lib.Masc[rand.Intn(len(lib.Masc))],
		lib.Neutral[rand.Intn(len(lib.Neutral))],
		lib.All[rand.Intn(len(lib.All))],
	)

	fmt.Println(out)
}

func pickmeup() {
	out := lib.GenderReduce(
		gender,
		"hey you, ma'am are amazing",
		"hey you, sir are amazing",
		"hey you, your are amazing",
		"hey you, your are amazing",
	)

	out += ", and your gender is valid. never forget how important you are."

	fmt.Println(out)
}
