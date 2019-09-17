package luhn

import(
  "strings"
  "strconv"
)

// Valid receives a number and uses Luhn checksum algorithm to determine its validity
func Valid(checkstr string) bool {
  checkstr = strings.Replace(checkstr," ","",-1)
  checkstrs := strings.Split(checkstr, "")

  // Strings less than 2 are auto-false
  if len(checkstrs) < 2 {
    return false
  }

  // Convert each string to integer
  numbers := make([]int, len(checkstrs))
  for i:=0; i<len(checkstrs); i++ {
    if _, err := strconv.Atoi(checkstrs[i]); err == nil {
      numbers[i], _ = strconv.Atoi(checkstrs[i])
    } else {
      return false
    }
  }

  tmpNum := 0
  sum := 0

  // Since Luhn counts every other digit backwards, we need to have a way to count
  // every other digit for even- or odd-length numbers
  checkI := 1
  if len(numbers) % 2 == 0 {
    checkI = 2
  }

  for i:=0; i<len(numbers); i++ {
    tmpNum = numbers[i]

    // Double every other digit counting backwards
    if checkI % 2 == 0 {
      tmpNum = tmpNum * 2

      if tmpNum > 9 {
        tmpNum = tmpNum - 9
      }
      numbers[i] = tmpNum
    }

    sum = sum + tmpNum
    checkI = checkI + 1
  }

  // Unless sum is divisible by 10, number is not valid checksum
  if sum % 10 > 0 {
    return false
  }

  return true
}
