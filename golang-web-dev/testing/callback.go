package main

import "fmt"


//callback 2
func filter(numbers []int, callback func(int) bool) []int {
  var xs []int
  for _, n := range numbers {
    if callback(n) {
      xs = append(xs, n)
    }
  }
  return xs
}

func main() {


  xr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  result := filter(xr, func(n int) bool {
    if (n % 2 == 0) {
      return true
    }
      return false
  })

  fmt.Println(xr)
  fmt.Println(result)
}
