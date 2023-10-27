package main
import "fmt" 
func main(){

	fmt.Println("VARIABLE SCOPE TUTORIAL")

	//city := "Faislabad"
	//fmt.Printf("The typr of city is %T",city)

	city := "Faislabad"
	
	//This is a seperate Block - inner block
	{
		country := "Pakistan"
		fmt.Println(city)
		fmt.Println(country)
	}
	
	fmt.Println(city)
	//fmt.Println(country)

}