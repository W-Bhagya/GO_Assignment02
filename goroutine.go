package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	result := make(chan int)

   var size1 int
   var size2 int

	fmt.Print("Enter size for chanel1 : ")
	fmt.Scanf("%d ",&size1)

	fmt.Print("Enter size of channel2 : ")
	fmt.Scanf("%d",&size2)

	// Channel 1
	go func() {
		for i := 0; i < size1; i++ {
			ch1 <- i + 1
			fmt.Println("Produced Channel1", i+1)
		}
		close(ch1)
	}()

	

    // Channel 2
	go func() {
		for i := 0; i < size2; i++ {
			ch2 <- i + 1
			fmt.Println("Produced Channel2", i+1)
		}
		close(ch2)
	}()

	go Produce(ch1, ch2, result)

	for x := range result {
		fmt.Println("Consumed Channel3", x)
	}

}
// Channel 3
func Produce(ch1, ch2, result chan int) {
	for x := range ch1 {
		result <- x
	}
	for x := range ch2 {
		result <- x
	}
	close(result)
}