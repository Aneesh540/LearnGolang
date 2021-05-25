# Basics of Maps

## 1. Initialisation
```golang
var map1 = map[string]int{}
map2 := make(map[string]int)
map3 := map[string]int{"bholu" : 1, "nitish" : 2}
var map4 map[string]int
```

## 2. Add & Delete
> It will delete if key is present otherwise pass silently
```golang
  map1["netflix"] = 123
  map1["prime"] = 456

  delete(map1, "netflixf")
```

## 3. Read operation
```golang
if value, ok := map1["netflix"]; ok { // check if value is presnt or not in map
		fmt.Println("netflix is present & value = ", value)
	}

ele, ok := map2["hulu"]
fmt.Println(ele, ok) // if ok = false, ele still contains default value of that data type
```

## 4. Iterate through Map
```golang
for key, value := range map3 {
		fmt.Println("===>>",key,value)
	}
```
