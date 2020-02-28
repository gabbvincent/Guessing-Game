//Name: Vincent Gabb
//Date: 2/27/2020
//Title: Guessing Game


package main

import (
  "fmt"
  "math/rand"
      

) 

func main() {
  
  var user int
  var computer = rand.Intn(101)

  fmt.Println("Enter a number between 1 and 100.")
  fmt.Scanln(&user)
  if user == computer {
    fmt.Println("Yay! You got it!")
  } else if user < computer {
    fmt.Println("Too low! Try again.")
    
    fmt.Scanln(&user)
  } else  {
    fmt.Println("Too high! Try again.")
    
    fmt.Scanln(&user)
  }

}
