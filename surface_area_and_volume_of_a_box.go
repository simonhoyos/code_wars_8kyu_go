package kata

func GetSize(w,h,d int) [2]int {
    return [2]int{w * h * 2 + w * d * 2 + h * d * 2, w * h * d}
}

