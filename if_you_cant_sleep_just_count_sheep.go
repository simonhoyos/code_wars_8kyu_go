package kata

import "fmt"

func countSheep(num int) string {
  var s string

  for i := 1; i <= num; i++ {
    s += fmt.Sprintf("%d sheep...", i)
  }

  return s
}

