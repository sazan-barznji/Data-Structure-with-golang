package main
import "fmt"


var slice []int  //declaring slice with out index 
func getdata(){
	var n int 
	var tempD int
	fmt.Println("enter how many data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("enter data: ")
		fmt.Scan(&tempD)	
		createslice(tempD) //appending value to slice 
	
	}
}
func createslice(x int )[] int{
	slice=append (slice, x)
	return slice
}

func main(){
	getdata()
	fmt.Println(slice) //output will be all slice numbers
}  
