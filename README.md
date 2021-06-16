# LearnGolang

## constants shadowing & untyped constants
## defer function call; pushed onto stack

## Strings & Rune
*Strings in Golang is different from other programming language*

*Strings are immutable*

*Strings are just collection of bytes, & bytes are alias of uint8 (unsigned int 8)*
#### Example 1
```golang
str := "aneesh Jain"
fmt.Printf("value=%v, Type=%T", str, str)
fmt.Printf("value=%v, Type=%T", str[2], str[2])

s[2] = 'z' // error
 ```
>**value=aneesh Jain, type=string**
>
>**value=101, type=uint8**

#### Example 2 (Strings & Byte interconversion, many times function expects byte as param)
```golang
 b := []byte("abcde")
 fmt.Printf("value=%v, Type=%T\n", b, b)

 s := string(b)
 fmt.Printf("value=%v, Type=%T\n", s, s)
```
> **value=[97 98 99 100 101], Type=[]uint8**
> 
> **value=abcde, Type=string**

#### Example 3 (Rune)
```golang
rune1 := '♄'   // remember single quotes
rune2 := '\a'
fmt.Printf("Rune 1: %c; Unicode: %U\n", rune1,rune1)
fmt.Printf("Rune 1: %c; Unicode: %U", rune2,rune2)
```
> **Rune 1: ♄; Unicode: U+2644**
>
> **Rune 1: ; Unicode: U+0007**



====================================================

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

===================================================================

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

## Defer function call

> Deferring a function executes it just before the function's return statement is executed