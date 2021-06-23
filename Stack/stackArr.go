package main
import "fmt"

var stack [10]int

func push(x int) {

	for i := 0; i < 10; i++ {
		if stack[i] == 0 {
			stack[i] = x
			return
		}
	}
}
// the functin can be called three time to 
func pop()  {
	var counter int 
	for (counter<3){		
		for i := len(stack)-1; i > -1; i-- {
			if stack[i] != 0 && stack[i]%2==0 {
				stack[i] = 0
				return
			}
		}	
	counter++
	}
}
func main(){
	push(1)
	push(2)
	push(3)
	push(4)
	push(5)
	push(6)

	fmt.Println(stack)
	pop()
	fmt.Println(stack)
	// pop()
	// fmt.Println(arr)
	// pop()
	// fmt.Println(arr)

}