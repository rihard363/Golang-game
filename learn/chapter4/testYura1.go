package main
import "fmt"
func Formation(ch chan<- int) { // Sequence for the channel
    for i := 2; ; i++ {
        ch <- i // We direct i into the ch channel.
    }
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
    for {
        i := <-in // Receive value from 'in'.
        if i%prime != 0 {
            out <- i // guide 'i' to 'out'.
        }
    }
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
    fmt.Print("Установите длину последовательности: ")
    var n int
    fmt.Scanf("%d", &n)
    ch := make(chan int) // Create a new channel.
    go Formation(ch)      // Launch Formation goroutine.
    for i := 0; i < n; i++ {
        prime := <-ch
        ch1 := make(chan int)
        go Filter(ch, ch1, prime)
        ch = ch1
    }
        p := <-ch
        fmt.Println(p)
}