package main

import "fmt"

func main() {
	var map1 = map[string]int{}
	map2 := make(map[string]int)
	map3 := map[string]int{"bholu" : 1, "nitish" : 2}
	var map4 map[string]int

	fmt.Println(map1, map2, map3, map4)

	// ============= CREATE OPERATIONS =============
	map1["netflix"] = 123
	map1["prime"] = 456


	// ============== DELETE OPERATIONS =============
	
	delete(map1, "netflixf")
	// delete value inside map => delete(map, key)
	// delete() doesn't return anything

	fmt.Println(map1)


	// ========== READ OPERATIONS ===============
	
	if value, ok := map1["netflix"]; ok { // check if value is presnt or not in map
		fmt.Println("netflix is present & value = ", value)
	}

	ele, ok := map2["hulu"]
	fmt.Println(ele, ok) // if ok = false, ele still contains default value of that data type

	if _, ok := map3["hotstar"] ; ok {
		fmt.Println("Hotstar is present")
	} else {
		fmt.Println("Hotstar NOT present")
	}

	// ========== ITERATING MAPS ============
	for key, value := range map3 {
		fmt.Println("===>>",key,value)
	}


}
