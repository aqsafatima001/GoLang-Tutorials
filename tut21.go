package main
import "fmt"

func main(){
	fmt.Println(SumOfNumbers())
	fmt.Println(SumOfNumbers(10))
	fmt.Println(SumOfNumbers(10, 20, 30, 40, 50))

	

}

func SumOfNumbers(num ...int) int {
	sum:=0
	for _, value := range num{

		sum += value
	}
	return sum
}