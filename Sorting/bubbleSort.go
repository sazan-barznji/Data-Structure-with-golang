package main
import "fmt"

var arr= [5]int {5,3,7,9,2}
var swapFound bool= true

func bubbleSort(){
	for j := 0; j < len(arr) &&swapFound; j++ {
		swapFound= false
		for i := 1; i < len(arr); i++ {
			if arr[i-1]>arr[i]{
				//swap them
				x:=arr[i-1]
				arr[i-1]=arr[i]
				arr[i]=x
				swapFound= true 
			}
		}
	}

}
func main(){
	bubbleSort()
	fmt.Println(arr)
}