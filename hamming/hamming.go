package hamming

import(
  "strings"
  "errors"
)

// Distance calculates Hamming Distance between two DNA strings
func Distance(a, b string) (int, error) {
  helixA := strings.Split(a, "")
  helixB := strings.Split(b, "")
  hammingDistance := int(0)

  if len(helixA) != len(helixB) {
    return hammingDistance, errors.New("Helix length mismatch")
  }

  for i:=0; i<len(helixA) ;i++ {
    if helixA[i] != helixB[i] {
      hammingDistance = hammingDistance + 1
    }
  }

  return hammingDistance, nil
}
