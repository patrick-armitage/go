package acronym

import(
	"strings"
)

// Abbreviate takes a string title and returns an uppercase acronym from its words
func Abbreviate(s string) string {
	s = strings.Replace(s,", "," ",-1)
	s = strings.Replace(s," - "," ",-1)
	s = strings.Replace(s,"-"," ",-1)
	s = strings.Replace(s,"_","",-1)

	acro := ""
	words := strings.Split(s, " ")
	upperChar := ""

	for i:=0; i<len(words); i++ {
		upperChar = strings.ToUpper(string(words[i][0]))
		acro = acro + upperChar
	}

	return acro
}
