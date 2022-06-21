package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type stack struct {
	st[100]int
	l int
	top int
}


// Push function
func push(s *stack, x int){

	if s.top >= s.l-1 {
		fmt.Println("Stack Overflow")
	} else {
		s.top++
		s.st[s.top] = x
	}

}
 
// Pop function
func pop(s *stack) int {
	var a int

	if s.top <= -1 {
		fmt.Println("Stack is empty")
		return -1
	} else { 
		a = s.st[s.top]
		s.st[s.top] = 0
		s.top--
	}
	return a
}


//Display function
func disp(s *stack) {

	for i := 0; i <= s.top; i++ {
		fmt.Println(s.st[i], "\t")
	}
}

func main(){
	rd := bufio.NewScanner(os.Stdin)
	s := stack{}
	fmt.Println("Enter size of stack")
	rd.Scan()
	val, _:= strconv.ParseInt(rd.Text(), 10,64)
	s.l = int(val)
	s.top = -1

	fmt.Println("*********Satck Operations********")
	fmt.Println("1. Push\n2. Pop\n3. Display\n4. Exit")

	for {
		fmt.Print("Enter Your Choice : ")
		rd.Scan()
		ch, _:= strconv.ParseInt(rd.Text(), 10,64)
		switch ch{
		
	    case 1 :
			fmt.Print("Enter The Element : ")
			rd.Scan()
			x, _:= strconv.ParseInt(rd.Text(), 10,64)
			push(&s, int(x))

		case 2 :
			b := pop(&s)
			if b > -1 {
				fmt.Println("Poped Element Is : ", b)
			}
		
		case 3 :
			fmt.Println("********Elements in Stack******")
			disp(&s)
			fmt.Println()
		
		case 4 :
			fmt.Println("Exit...")
			return
		
		default :
			fmt.Println("Wrong Choice...")

		}	
	}
}

