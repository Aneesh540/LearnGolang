package main
import (
	"fmt"
)

type bot interface {
	getGreetings() string 
	// this tell that : those structs who have getGreetings() function as their property and return string
	// then automatically they are also the honorary member of "bot" interface type
	// Now you are member of "bot" interface type, you can use printGreeting() function also
	// u cannot create value of interface directly

}

func printGreeting(b bot){
	fmt.Println(b.getGreetings())
}

type englishBot struct {}
type spanishBot struct {}

func (eb englishBot) getGreetings() string {
	return "Hello"
}
func (sb spanishBot) getGreetings() string {
	return "Holaaa!"
}

func main() {

	english := englishBot{}
	spanish := spanishBot{}

	printGreeting(english)
	printGreeting(spanish)

}
