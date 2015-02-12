package main

import "fmt"

// let's define a Person
type Person struct {
    Name string
}

// who's able to introduce herself
func (p *Person) Introduce() {
    fmt.Printf("Hola, me llamo %s.\n", p.Name)
}

// composition: one struct into another, letting it access its methods
type Saiyan struct {
    *Person
    Power int
}


func main()  {

    // un goku
    persona := &Person{"Manolo"}

    persona.Introduce()
    fmt.Println(persona)

    // un goku
    goku := &Saiyan{
        Person: &Person{"Goku"},
        Power: 9000,
    }
    // un krillin
    krillin := Saiyan{
        Person: &Person{"Krillin"},
        Power: 2000,
    }

    goku.Power += 100
    goku.Introduce()
    fmt.Println(goku)
    fmt.Printf("%s %d\n", goku.Name, goku.Power)

    krillin.Power += 100
    krillin.Introduce()
    fmt.Println(krillin)
    fmt.Printf("%s %d\n", krillin.Name, krillin.Power)

}
