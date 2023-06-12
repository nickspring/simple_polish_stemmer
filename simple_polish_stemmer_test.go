package simple_polish_stemmer

import (
	"strings"
	"testing"
)

const (
	accentInputText      = "Kariera na językach to wydarzenie zorganizowane z myślą o studentach i absolwentach znających języki obce na poziomie co najmniej Będą oni mieli okazję zastanowić się nad kierunkami rozwoju własnej kariery zawodowej w oparciu o informacje na temat możliwości wykorzystania swoich umiejętności lingwistycznych na współczesnym rynku pracy dlatego też nie chcę"
	accentExpectedText   = "karier na język to wydarz zorganizowane z myśl o studen i absolwen znaj język obce na poziom co najmn będą oni miel okazj zastanow się nad kierunk rozwoj własn karier zawodow w opar o informacje na temat możliwośc wykorzys swoich umiejętnośc lingwistyczn na współczesnym rynk prac dlat też nie chcę"
	unaccentExpectedText = "karier na jezyk to wydarz zorganizowan z mysl o studen i absolwen zna jezyk obce na poziom co najmn beda oni miel okazj zastan sie nad kierunk rozwoj wlasn karier zawod w opar o informacj na temat mozliwosc wykorzys swoich umiejetnosc lingwistyczn na wspolczesnym rynk prac dlat tez nie chce"
)

func TestSimplePolishStemmer(t *testing.T) {

	words := strings.Split(accentInputText, " ")
	expectedWords := strings.Split(accentExpectedText, " ")
	var unstemmableList []string

	for index, word := range words {
		expectedWord := expectedWords[index]
		resultWord := StemWord(word, false, unstemmableList)
		if resultWord != expectedWord {
			t.Errorf("'%s' != '%s'", resultWord, expectedWord)
		}
	}

}

func TestSimplePolishStemmerUnaccented(t *testing.T) {

	words := strings.Split(accentInputText, " ")
	expectedWords := strings.Split(unaccentExpectedText, " ")
	var unstemmableList []string

	for index, word := range words {
		expectedWord := expectedWords[index]
		resultWord := StemWord(word, true, unstemmableList)
		if resultWord != expectedWord {
			t.Errorf("'%s' != '%s'", resultWord, expectedWord)
		}
	}

}
