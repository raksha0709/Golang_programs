package main

import "fmt"

func main() {
	mymap := make(map[string]int)
	var key string
	var value int
	for {
		fmt.Println("Enter a key(or enter quit):")
		fmt.Scan(&key)
		if key == "quit" {
			break
		}
		fmt.Println("Enter the value")
		fmt.Scan(&value)
		mymap[key] = value
	}
	fmt.Println("Map items are:", mymap)
	//accessing
	fmt.Println("enter the key to be accessed:")
	fmt.Scan(&key)
	if value, ok := mymap[key]; ok {
		fmt.Printf("The value at key %s is:%d\n", key, value)
	} else {
		fmt.Println("The key is not found!...")
	}
	//deleting
	fmt.Println("enter the key to be deleted:")
	fmt.Scan(&key)
	delete(mymap, key)
	fmt.Println("After deletion the map is", mymap)
	//find lenght
	fmt.Println("The lenght of the map is:", len(mymap))
}
