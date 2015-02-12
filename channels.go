package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    //c := make(chan int)
    // buffer channel
     c := make(chan int, 100)
    for index := 0;  index < 5; index++ {
        worker := &Worker{id: 1}
        go worker.process(c)
    }

    for {
        c <- rand.Int()
        time.Sleep(time.Millisecond * 50)
    }
}


type Worker struct {
    id int
}

func (w *Worker) process (c chan int) {
    for {
        data := <- c
        fmt.Printf("worker %d got %d\n", w.id, data)
        fmt.Println(len(c))
        time.Sleep(time.Millisecond * 500)
    }
}
