package kata

import "math"

func ExpressionMatter(a int, b int, c int) int {
  var temp float64

  res := [...]int{a*b*c, a*b+c, a+b*c, a+b+c, (a+b)*c, a*(b+c)}

  for _,v := range res {
    temp = math.Max(temp, float64(v))
  }

  return int(temp)
}

