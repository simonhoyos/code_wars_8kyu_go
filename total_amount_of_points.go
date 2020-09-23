package kata

import "strconv"

func Points(games []string) int {
  s := 0

  for _, r := range games {
    v1, _ := strconv.Atoi(string(r[0]))
    v2, _ := strconv.Atoi(string(r[2]))

    if v1 > v2 {
      s = s + 3
    } else if v2 > v1 {
      s = s + 0
    } else {
      s = s + 1
    }
  }

  return s;
}

