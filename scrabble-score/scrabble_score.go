package scrabble

import(
  "strings"
)

// Score takes a Scrabble word and calculates the total score for that word
func Score(word string) int {
  pointsMap := make(map[string]int)
  pointsMap["AEIOULNRST"] = 1
  pointsMap["DG"] = 2
  pointsMap["BCMP"] = 3
  pointsMap["FHVWY"] = 4
  pointsMap["K"] = 5
  pointsMap["JX"] = 8
  pointsMap["QZ"] = 10

  score := 0
  letters := strings.Split(word, "")
  currLetter := ""
  for i:=0; i<len(letters); i++ {
    for keyletters, points := range pointsMap {
      currLetter = strings.ToUpper(letters[i])

      if strings.Contains(keyletters, currLetter) {
        score = score + points
      }
    }
  }

  return score
}
