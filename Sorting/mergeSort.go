
package main
import "fmt"

var slice []int 
//merge sort 
func mergesort(input []int) []int {
      
    if len(input)<2 {
        return input
	}
	mid := int(len(input)/2)  
	input= merge(mergesort(input[:mid]), mergesort(input[mid:]) )
	return input
}

//merge descending order

func merge(left []int , right []int) []int {
	size := len(left) + len(right)
	slice= make([]int, size,size)
      
	i := 0
	j:=0

	for k := 0; k < size; k++ {
		
        if i> len(left)-1  && j <= len(right)-1 {
			slice[k] = right[j]
			j++
        } else if j>len(right)-1 && i<= len(left)-1{
            slice[k] = left[i]
           i++
        } else if left[i]> right [j]{
			slice[k]= left[i]
			i++
		}else {
			slice [k]= right[j] 
			j++
		}
        
    }
	fmt.Println("in the merge ",slice)
    return slice
}


//merge  ascending order 
/*    
 
func merge(left []int , right []int) []int {
	size := len(left) + len(right)
	slice= make([]int, size,size)
      
	i := 0
	j:=0

	for k := 0; k < size; k++ {
		
        if i> len(left)-1  && j <= len(right)-1 {
			slice[k] = right[j]
			j++
        } else if j>len(right)-1 && i<= len(left)-1{
            slice[k] = left[i]
           i++
        } else if left[i]< right [j]{
			slice[k]= left[i]
			i++
		}else {
			slice [k]= right[j] 
			j++
		}
        
    }
	fmt.Println("in the merge ",slice)
    return slice
}
 ;*/