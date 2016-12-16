package main

import "fmt"

func main() {
	//fan out
	//distribute the sq work across two goroutines that both read from in.
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	//fan in
	//consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return output
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg.sync.WaitGroup
    
    // start an output goroutine for each input channel in cs
    //output copies values from c to out until c is closed, then call wg.Done
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }

    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()
    return out

}
