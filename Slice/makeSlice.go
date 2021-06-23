package main
import "fmt"

var slice [] int
var index int 

func makeSlice(grade int) []int {	
	slice[index] = grade
	index++       
	return slice 
}

func getdata(){
	var n int 
	var tempD int
	
	fmt.Println("enter how many data: ")
	fmt.Scan(&n)

	slice= make([]int,n,n)

	for i := 0; i < n; i++ {
		fmt.Println("enter data: ")
		fmt.Scan(&tempD)	
		makeSlice(tempD) 
	
	}
}
func main(){
	getdata()
	fmt.Println(slice)
}