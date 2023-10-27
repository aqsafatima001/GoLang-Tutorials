package main
import "fmt"
func main() {

	var grades [5] int
	grades = [5] int {98,80,78,54,24}
	fmt.Println(grades)

	var fruits [5] string
	fruits = [5] string {"mango", "apple" , "Banana" , "Grapes", "Orange"}
	fmt.Println(fruits)

	names := [...] string {"Aqsa", "Reha" , "Rimsha" , "shaheer" , "haseeb" , "Zaid"}
	fmt.Println(names)
	fmt.Println()
	fmt.Println("The Length of names[] is = ",len(names))
	fmt.Println("The value at first Index of names[] is = ", names[0])
	fmt.Println()
	fmt.Println()

	fmt.Println(names)
	names[3] = "Shaheer Siraj"
	fmt.Println(names)


	// for i:=0; i<len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index , element := range names{
		fmt.Println(index, "=>", element)
	}

}
