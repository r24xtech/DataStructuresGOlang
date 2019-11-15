package main

import "fmt"

func main() {
  var numbers []int = []int{5, 4, 2, 3, 1, 0}
  fmt.Println("Our list of numbers is: ", numbers)
  bubbleSort(numbers)
  fmt.Println("After sorting: ", numbers)
}

func bubbleSort(numbers []int) {
  var N int = len(numbers)
  var i int
  for i = 0; i < N; i++ {
    if !sweep(numbers, i) {
      return
    }
  }
}

func sweep(numbers []int, prevPasses int) bool{
  var N int = len(numbers)
  var firstIndex int = 0
  var secondIndex int = 1
  var didSwap bool = false

  for secondIndex < (N - prevPasses) {
    var firstNumber int = numbers[firstIndex]
    var secondNumber int = numbers[secondIndex]
    if firstNumber > secondNumber {
      numbers[firstIndex] = secondNumber
      numbers[secondIndex] = firstNumber
      didSwap = true
    }

    firstIndex++
    secondIndex++
  }
  return didSwap
}
