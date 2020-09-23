package kata

import "strings"

func Solution(word string) string {
  s := strings.Split(word, "")

  var a string
  for _, r := range s {
    a = string(r) + a
  }

  return a
}

