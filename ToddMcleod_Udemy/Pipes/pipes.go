package main

func main() {
	cWriter := make(chan int)
	go func() {
		cWriter <- 2
	}()

	cReader := <-cWriter
	println(cReader)
}
