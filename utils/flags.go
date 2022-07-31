package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/bartick/AniLang/errors"
)

type FlagStruct struct {
	Name    string
	Value   string
	Help    string
	Boolean bool
}

var Flags = make(map[string]FlagStruct)
var (
	FilePath   string = ""
	flagPassed bool   = false
)

func AddFlag(flagName string, flagHelp string) {
	Flags[flagName] = FlagStruct{
		Name:    flagName,
		Value:   "",
		Help:    flagHelp,
		Boolean: false,
	}
}

func AddBoolFlag(flagName string, flagHelp string) {
	Flags[flagName] = FlagStruct{
		Name:    flagName,
		Value:   "true",
		Help:    flagHelp,
		Boolean: true,
	}
}

func validateFlags(flagToCheck string) (int, bool) {
	if strings.HasPrefix(flagToCheck, "--") && len(flagToCheck) > 2 {
		return 2, true
	} else if strings.HasPrefix(flagToCheck, "-") && len(flagToCheck) > 1 {
		return 1, true
	} else {
		return 0, false
	}
}

func FlagParser() {

	AddBoolFlag("help", "Width of the animation")

	for i := 1; i < len(os.Args); i++ {

		if !flagPassed {
			flagPassed = true
		}

		if delemeter, flagValid := validateFlags(os.Args[i]); flagValid {
			if entry, isPresent := Flags[os.Args[i][delemeter:]]; isPresent {

				if entry.Boolean {
					entry.Value = "true"
					Flags[os.Args[i][delemeter:]] = entry
				} else if i+1 < len(os.Args) {
					entry.Value = os.Args[i+1]
					Flags[os.Args[i][delemeter:]] = entry
					i++
				} else {
					// TODO: Throw error
				}
			}
			continue
		} else if strings.HasSuffix(os.Args[i], ".ani") {

			if FilePath != "" {
				// TODO: throw error
			}

			FilePath = os.Args[i]
		} else {

			fmt.Println(errors.Errors[errors.NVF])

			os.Exit(1)
		}

	}

	if flagPassed && FilePath == "" {

		fmt.Println(errors.Errors[errors.NAF])

		os.Exit(1)
	}

}
