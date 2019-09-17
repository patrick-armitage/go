package diffsquares

import()

// NaturalNums returns all natural numbers leading up to and including num
func NaturalNums(num int) []int {
  numbers := make([]int,0)
  for i:=0; i<num; i++ {
    numbers = append(numbers, i+1)
  }

  return numbers
}

// SquareOfSum takes a list of numbers and returns the square of their sum
func SquareOfSum(num int) (square int) {
  numbers := NaturalNums(num)
  sum := 0
  for i:=0; i<len(numbers); i++ {
    sum = sum + numbers[i]
  }

  square = sum * sum
  return square
}

// SumOfSquares takes a list of numbers and returns the sum of their squares
func SumOfSquares(num int) (sum int) {
  numbers := NaturalNums(num)
  tmpNum := 0
  for i:=0; i<len(numbers); i++ {
    tmpNum = numbers[i] * numbers[i]
    sum = sum + tmpNum
  }

  return sum
}

// Difference takes a list of numbers and returns SquareOfSum - SumOfSquares
func Difference(num int) (diff int) {
  squareOfSum := SquareOfSum(num)
  sumOfSquares := SumOfSquares(num)

  diff = squareOfSum - sumOfSquares
  return diff
}
