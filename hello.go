package main

import (
  "fmt"
  "math"
  "rsc.io/quote"
)

var(
  isHappy   bool  = true
  shortint  uint8 = 2
  str       string= "GM!"
)


func add(x int, y int) int {
  return x + y
}

// Something interesting
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

type timestamps struct {
  Begin [2] float64
  End [2] float64
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func (t timestamps) min() float64 {
	var smallest float64
	for v := range t.Begin {
		if smallest == 0 {
			smallest = t.Begin[v] 
		}
		smallest = math.Min(smallest, t.Begin[v])
	}
	return smallest
}


func main() {
  fmt.Println(quote.Go())
  // assign output of add() to var s using short declare method :=
  s := add(40,2)
  // declare integer sum and assign output to it
  var sum int = add(6,9)
  
  fmt.Println(s)
  fmt.Println(sum)

  fmt.Printf("Type: %T, Value: %v\n", isHappy, isHappy)
  fmt.Printf("Type: %T, Value: %v\n", shortint, shortint)
  fmt.Printf("Type: %T, Value: %v\n", str, str)

  // Type convertion
  a := 5.63
  b := fmt.Sprintf("%f",a)
  c := uint8(a)

  fmt.Println(a)  
  fmt.Println(b)
  fmt.Println(c)  

  // for loop
  for i:=0; i<3; i++ {
    fmt.Println(i)
  }

  // While loop °_°'
  total := 0
  for total<5 {
    total++
  } 
  
  // If statement
  if total>3 {
    fmt.Println("Big success")
  }

  // Some strange func with for and if convertion
  e := pow(3,2,10)
  fmt.Println(e)

  // A defer statement defers the execution of a function until the surrounding function returns.
  defer fmt.Println("This is a deferred call")
  fmt.Println("This call is not deferred")

	// Pointers
	f := 42
	p := &f
  fmt.Printf("Value of pointer: %v\n", p)
  fmt.Println("The actual value:", *p)
  *p = 84
  fmt.Println("New value for original:", f)

  // Structures
  ts := timestamps{}
  ts.Begin = [2]float64{5.31,5.45}
  ts.End = [2]float64{8.72, 8.8}
  fmt.Println(ts.End[1])

  // Maps (key, value pairs)
  m := make(map[string]uint8)
  m["age"]=20
  fmt.Println("age:", m["age"])

  var prices = map[string]float64{
    "cookie"    : 1.50,
    "milkshake" : 2.20,
  }
  fmt.Println("Milkshake price:", prices["milkshake"])

  // Mutating maps
  delete(prices, "cookie")
  if _,ok := prices["cookie"]; !ok {
    fmt.Println("Cookie price removed!")
  }

	// Some more interesting stuff
  pos, neg := adder(), adder()
  for i := 0; i < 3; i++ {
    fmt.Println(
      pos(i),
	    neg(-2*i),
    )
  }

  // Methods (go doesn't have classes, but you can make methods for types)
  fmt.Println(ts.min())

}

