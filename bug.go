func myFunc(a []int) []int {
  b := make([]int, len(a))
  for i := range a {
    b[i] = a[i] * 2
  }
  return b
}

func main() {
  x := []int{1, 2, 3, 4, 5}
  y := myFunc(x)
  fmt.Println(y) // Output: [2 4 6 8 10]

  x[0] = 100
  fmt.Println(x) // Output: [100 2 3 4 5]
  fmt.Println(y) // Output: [2 4 6 8 10] // Unexpected behavior if you expect y to change
}