package main
import "fmt"

func main(){

	var a int
	var b int

	fmt.Println("Enter the value of a: ")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Printf("Error reading input for a: %v\n", err)
		return
	}

	fmt.Scanln()

	fmt.Println("Enter the value of b: ")
	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		fmt.Printf("Error reading input for b: %v\n", err)
		return
	}

	fmt.Println("The sum of a+b = ", AddNumbers(a,b))
}

func AddNumbers(a int , b int) int {
	sum:=a+b
	return sum
}