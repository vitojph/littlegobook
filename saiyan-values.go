package main

import "fmt"

// a Saiyan struct
type Saiyan struct {
    Name string
    Power int
}

// a method for the Saiyan struct
func (s Saiyan) Super() {
    s.Power += 10000
}


func main() {

    goku := Saiyan{"Goku", 9000}
    goku.Super()
    fmt.Printf("%s tiene %d de fuerza\n", goku.Name, goku.Power)

}
