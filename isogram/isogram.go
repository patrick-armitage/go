package isogram

import(
  "strings"
)

// IsIsogram takes a string and returns a boolean whether it is or is not an isogram
func IsIsogram(word string) bool {
  upcaseWord := strings.ToUpper(word)
  upcaseWord = strings.Replace(upcaseWord,"-","",-1)
  upcaseWord = strings.Replace(upcaseWord," ","",-1)
  letters := strings.Split(upcaseWord, "")
  currLetter := ""
  tmpString := ""

  for i:=0; i<len(letters); i++ {
    currLetter = letters[i]
    tmpString = strings.Replace(upcaseWord,currLetter,"",1)
    if strings.Contains(tmpString, currLetter) {
      return false
    }
  }

  return true
}
