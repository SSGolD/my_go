package main 
import "fmt"

 func main() {
	//for
	for i := 0; i < 5; i++ {
		fmt.Println("SerJ")
	}
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	s := 0 // sum 1 to 100 
	for i := 0; i < 101; i++ {
	    s += i;
	}
	fmt.Println(s)
 }
 
