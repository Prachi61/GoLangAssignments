// _Interfaces_ are named collections of method
// signatures.

package main

import "fmt"


// Here's a basic interface for geometric shapes.
type planet interface {
    name() string
    mass() int
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type planets struct {
    pname string
 pmass int
}


// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (p planets) name() string{
    return p.pname
}
func (p planets) mass() int{
    return p.pmass
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func details(g planet) {
    fmt.Println(g)
    fmt.Println(g.name())
    fmt.Println(g.mass())
}

func main() {
  mercury := planets{pname:"mercury", pmass:4}
	details(mercury)

	venus := planets{pname:"venus", pmass:5}
	details(venus)

	earth := planets{pname:"earth", pmass:6}
	details(earth)
	
	mars := planets{pname:"mars", pmass:8}
	details(mars)

   
}