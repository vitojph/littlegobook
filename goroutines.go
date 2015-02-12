package main

import (
    "fmt"
    "time"
)

func process()  {
    fmt.Println("processing")
}

func main() {

    fmt.Println("start")
    go process()
    time.Sleep(time.Millisecond * 50)
    fmt.Println("done")

}
