func myFunc(a []int) []int {
  b := make([]int, len(a))
  copy(b, a) // Correct way to create a copy of the slice
  for i := range b {
    b[i] = b[i] * 2
  }
  return b
}

func main() {
  x := []int{1, 2, 3, 4, 5}
  y := myFunc(x)
  fmt.Println(y) // Output: [2 4 6 8 10]

  x[0] = 100
  fmt.Println(x) // Output: [100 2 3 4 5]
  fmt.Println(y) // Output: [2 4 6 8 10]
} 