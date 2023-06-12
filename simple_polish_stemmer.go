package simple_polish_stemmer

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
)

/* Exported functions & variables */

// StemWord stem one word, unaccentedMode produces unaccented stems
func StemWord(word string, unaccentedMode bool, unstemmableWords []string) string {

	// prepare word
	word = strings.ToLower(word)
	if unaccentedMode {
		word = unaccent(word)
	}

	// prepare unstemmableWords
	for _, uWord := range unstemmableWords {
		uWord = strings.ToLower(uWord)
		if unaccentedMode {
			uWord = unaccent(uWord)
		}
		if word == uWord {
			return word
		}
	}

	for _, rule := range stemmerRules {
		run := []rune(word)
		if len(run) < rule.MinWordLen {
			continue
		}

		// get suffixes
		suffixes := rule.SuffixesAccented
		if unaccentedMode && len(rule.SuffixesUnaccented) > 0 {
			suffixes = rule.SuffixesUnaccented
		}
		// check suffixes
		hasSuffix := false
		for _, suffix := range suffixes {
			if hasSuffix = strings.HasSuffix(word, suffix); hasSuffix {
				break
			}
		}
		// process according to the rule
		if hasSuffix {
			word = string(run[rule.LeftShift : len(run)-rule.RightShift])
		}
	}
	return word
}

/* Internal functions & variables */

// stemmerRule describes one stemmer rule
type stemmerRule struct {
	MinWordLen       int
	LeftShift        int
	RightShift       int
	SuffixesAccented []string
	// empty if the same
	SuffixesUnaccented []string
}

// stemmerRules describes stemmer rules and their consequence
var stemmerRules = []stemmerRule{

	// Remove nouns
	{8, 0, 4,
		[]string{"zacja", "zacją", "zacji"},
		[]string{},
	},
	{7, 0, 4,
		[]string{"acja", "acji", "acją", "tach", "anie", "enie", "eniu", "aniu"},
		[]string{},
	},
	{7, 0, 2,
		[]string{"tyka"},
		[]string{},
	},
	{6, 0, 3,
		[]string{"ach", "ami", "nia", "niu", "cia", "ciu"},
		[]string{},
	},
	{6, 0, 2,
		[]string{"cji", "cja", "cją"},
		[]string{},
	},
	{6, 0, 2,
		[]string{"ce", "ta"},
		[]string{},
	},

	// Diminutive
	{7, 0, 5,
		[]string{"eczek", "iczek", "iszek", "aszek", "uszek"},
		[]string{},
	},
	{7, 0, 2,
		[]string{"enek", "ejek", "erek"},
		[]string{},
	},
	{5, 0, 2,
		[]string{"ek", "ak"},
		[]string{},
	},

	// Remove adjectives ends
	{8, 3, 3,
		[]string{"naj", "sze", "szy"},
		[]string{},
	},
	{8, 3, 5,
		[]string{"naj", "szych"},
		[]string{},
	},
	{7, 0, 4,
		[]string{"czny"},
		[]string{},
	},
	{6, 0, 3,
		[]string{"owy", "owa", "owe", "ych", "ego"},
		[]string{},
	},
	{6, 0, 2,
		[]string{"ej"},
		[]string{},
	},

	// Remove verbs ends
	{6, 0, 3,
		[]string{"bym"},
		[]string{},
	},
	{6, 0, 3,
		[]string{"esz", "asz", "cie", "eść", "aść", "łem", "amy", "emy"},
		[]string{"esz", "asz", "cie", "esc", "asc", "lem", "amy", "emy"},
	},
	{4, 0, 2,
		[]string{"esz", "asz", "eść", "aść", "eć", "ać"},
		[]string{"esz", "asz", "esc", "asc", "ec", "ac"},
	},
	{4, 0, 1,
		[]string{"aj"},
		[]string{},
	},
	{4, 0, 2,
		[]string{"ać", "em", "am", "ał", "ił", "ić", "ąc"},
		[]string{"ac", "em", "am", "al", "il", "ic", "ac"},
	},

	// Remove adverbs ends
	{5, 0, 2,
		[]string{"nie", "wie"},
		[]string{},
	},
	{5, 0, 2,
		[]string{"rze"},
		[]string{},
	},

	// Remove plural forms
	{5, 0, 2,
		[]string{"ów", "om"},
		[]string{"ow", "om"},
	},
	{5, 0, 3,
		[]string{"ami"},
		[]string{},
	},

	// General ends
	{5, 0, 2,
		[]string{"ia", "ie"},
		[]string{},
	},
	{5, 0, 1,
		[]string{"u", "ą", "i", "a", "ę", "y", "ę", "ł"},
		[]string{"u", "a", "i", "e", "y", "l"},
	},
}

// unaccent replaces diacritics with appropriate ASCII symbols
func unaccent(s string) string {
	t := transform.Chain(
		norm.NFD,
		precis.UsernameCaseMapped.NewTransformer(),
		runes.Map(func(r rune) rune {
			switch r {
			case 'ą':
				return 'a'
			case 'ć':
				return 'c'
			case 'ę':
				return 'e'
			case 'ł':
				return 'l'
			case 'ń':
				return 'n'
			case 'ó':
				return 'o'
			case 'ś':
				return 's'
			case 'ż':
				return 'z'
			case 'ź':
				return 'z'
			}
			return r
		}),
		norm.NFC,
	)
	output, _, _ := transform.String(t, s)
	return output
}
