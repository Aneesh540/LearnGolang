# LearnGolang

## Shadowing of variables in block scope
## Strings : different from others programming languages
## constants shadowing & untyped constants
## defer function call; pushed onto stack

## Functions with multiple *args

```golang
func print(args ...int) int {
  total := 0
  for _, v := range args {
  total += v
  }
  return total
}

func main() {
  fmt.Println(add(1,2,3))
  fmt.Println(add(1,2,3,4,5))
}
```

## Pointers (GOLANG IS ALWAYS PASS BY VALUE)
> Pointers use for general data type
```golang
func change(stringPtr *string){
    *stringPtr = "bholu"
}

func main() {
   x := "bhai"
   Ptr := &x
   change(Ptr)
   fmt.Println(x) // bholu
   
}
```

> Pointers #2, For structs shortcut 
> We can use above method for structs, but here is a shortcut
> Whenever a receiver function is like this, go will allow us to call 
> the method on pointer to struct 
```golang
type person struct {
    name string
    salary int
}

func (PtrToPerson *person) updateName(newName string){
    (*PtrToPerson).name = newName
}

func main() {
   x := person{"aneesh", 100}
   x.updateName("nitish")
   fmt.Println(x) // {nitish 100}
   
}
```
