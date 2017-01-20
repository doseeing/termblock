package main

import (
    "math/rand"
    "time"
    "fmt"

    "github.com/doseeing/termblock"
)

func main() {

    tb := termblock.New(`FT`)
    tb.Print([]int{0,1,0,1,0,1}, "Status: ")

    show(tb, "Status LED", `▌`, []int{31, 32, 33}, 3)

    show(tb, "Blocks", termblock.CharsetBlocks, []int{32, 36}, 10)

    show(tb, "LOVE", `ILOVEU`, []int{31,32,33,34,35,36}, 6)

    show(tb, "Dots", termblock.CharsetDots, []int{31,32,33,34,35,36}, 100)

    show(tb, "Directions", termblock.CharsetDirections, []int{32}, 4)
}

func show(tb *termblock.TermBlock, name, alphabet string, color []int, randRange int) {
    tb.SetColorAlphabet(alphabet, color)

    fmt.Printf("\n\n  %s: %s\n\n", name, alphabet)

    for i := 0; i < 30; i++ {
        
        status := []int{}
        for j:=0; j<30; j++ {
            status = append(status, rand.Intn(randRange))
        }

        tb.Print(status, `Status:`)

        time.Sleep(100 * time.Millisecond)
    }
}