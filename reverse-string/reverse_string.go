package reverse

import(

)

// Reverse takes a string and returns it as a reversed mirror string
func Reverse(word string) (result string) {
  for _,v := range word {
    result = string(v) + result
  }

  return
}
