package main

import "fmt"

func main()  {
  // var can be used inside or outside of a function
  // variable declaration and value assigment can be done separeately
  var student1 string = "Rangga"
  var student2 = "Falah"
  var s string
  var i int
  var bo bool

  // := can only be used inside of a function
  // variable declaration and value assigment cannot be done separately
  x := 2

  // You can also declared multiple variables in the same line
  var a, b, c, d int = 1,2,3,4

  // You can also declared variable in block
  var (
    z int 
    u int = 1
    m string = "Miftahul"
  )

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
  fmt.Println(s)
  fmt.Println(i)
  fmt.Println(bo)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(z)
  fmt.Println(u)
  fmt.Println(m)
}
