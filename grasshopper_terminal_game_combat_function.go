package kata

import "math"

func combat(health, damage float64) float64 {
  return math.Max(0.0, health - damage)
}

