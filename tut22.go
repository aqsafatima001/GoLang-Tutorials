package main
import "fmt"

func main(){
	var n int 
	fmt.Print("Enter the Number you want to take the factorial of : ")
	fmt.Scanf("%d" , &n)
	fmt.Scanln()

	result := factorial(n)

	fmt.Printf("The Factorial of %d  =  %d" , n, result)
}

func factorial (n int )int{

	if n==0 {
		return 1
	}else{
		return n * factorial(n-1)
	}
}