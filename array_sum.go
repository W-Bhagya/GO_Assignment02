package main

import "fmt"

func routine_1 (arr []int, ch chan int) {

	sum := 0
    
	for _, n := range arr {
		sum += n
	}

	ch <- sum
}


func routine_2 (arr []int, ch chan int) {

	sum := 0
    
	for _, n := range arr {
		sum += n
	}

	ch <- sum
}


func array(x, y int) []int {
	arr := []int{}

	for i := x; i <= y; i++ {
		arr = append(arr,i)
	}

	return arr
}


func main() {

	var (
		st1 int
		st2 int
	    end1 int
		end2 int
	)

	fmt.Print("Enter starting point for routine1 & routine2 : ")
	fmt.Scanf("%d %d ",&st1,&st2)

	fmt.Print("Enter ending point for routine1 & routine : ")
	fmt.Scanf("%d %d",&end1,&end2)

	r1 := array(st1,end1)
	r2 := array(st2,end2)
	result := make(chan int)

	go routine_1(r1,result)
	go routine_2(r2,result)

	fmt.Println("Sum = ",<-result + <-result)
}

