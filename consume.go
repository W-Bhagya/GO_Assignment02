package main

import "fmt"

func main() {

	var size int
	fmt.Print("Enter size of array : ")
	fmt.Scanf("%d",&size)
	
     ch := consumeChannel() 

		for i := 0; i <= size; i++{
		ch <- i+1
		fmt.Println("Produced",i+1)
	}
}
func consumeChannel() chan <- int {

	ch := make(chan int)

	go func() {
		for i:= range ch {
			fmt.Println("Consumed",i)
		
		}
	}()

	return ch
}