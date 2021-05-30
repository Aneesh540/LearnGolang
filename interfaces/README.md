## Golang Interfaces

#### Normal Code
```golang
type englishBot struct {}
type spanishBot struct {}

func (eb englishBot) getGreetings() string {
	return "Hello"
}
func (sb spanishBot) getGreetings() string {
	return "Holaaa!"
}
```
#### Interface define

```golang
type bot interface {
	getGreetings() string 
 }
 
 func printGreeting(b bot){
	fmt.Println(b.getGreetings())
}
```

 1. this interface is telling us that
    1. Listen all those structs who have getGreetings() function as their property and return string
    2. NOW you all are automatically the honorary member of "bot" interface type
    3. As a member of "bot" interface type, you can now use printGreeting() function without any errors

