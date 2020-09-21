package kata

func century(year int) int {
 if rem := year % 100; rem == 0 {
   return year / 100
 } else {
   return (year / 100) + 1
 }
}

