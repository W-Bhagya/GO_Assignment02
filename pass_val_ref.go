package main

import "fmt"

type update struct {
	pages int

}

// Pass by value
func new_pages(u update ) {

	fmt.Println("Total no. of pages in a book Before :  ",u)
	u.pages++
	fmt.Println("Total no. of pages in a book After : ",u)

}

// Pass by reference
func pages_added(u *update )  {

	fmt.Println("Number of pages in a book Before :  ",u)
	u.pages++
	fmt.Println("Number of pages in a book After : ",u)
	

}

func main() {

	x := update {
		pages : 546,

	}


	fmt.Println("Number of pages in a book : ",x.pages)
	
	new_pages(x)
	fmt.Println("\n*********AFFTER PASS BY VALUE*********")
	fmt.Println("Number of pages in a book : ",x.pages)

	pages_added(&x)
	fmt.Println("\n*********AFFTER PASS BY REFERENCE*********")
	fmt.Println("Number of pages in a book : ",x.pages)

	fmt.Println("Total Number Of Pages In A Book : ",x.pages)


}