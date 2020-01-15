package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"sort"
)

func main() {
	//声明map 未初始化
	var dictionary map[string]int
	dictionary = make(map[string]int)

	dictionary["demo"] = 100
	dictionary["alpha"] = 200
	//fmt.Println(dictionary)

	value, ok := dictionary["demo"]
	if ok {
		fmt.Println("value = ", value)
	} else {
		fmt.Println("Nothing")
	}

	fmt.Println("Loop of the map : ")

	for value, key := range dictionary {
		fmt.Println(value, "	", key)
	}

	fmt.Println("delete a element in map : ")
	dictionary["deleted"] = 0
	delete(dictionary, "deleted")
	fmt.Println(dictionary)

	sortKeywithMap()

}

func sortKeywithMap() {
	rand.Seed(time.Now().UnixNano())

	var dictionary = make(map[string]int, 10)

	for i := 0; i < 10; i++ {
		key := "s" + strconv.Itoa(rand.Intn(100))
		dictionary[key] = i

	}
	var listofKey = make([]string, 0, len(dictionary))
	for key := range dictionary {
		listofKey = append(listofKey, key)
	}


	sort.Strings(listofKey)
	
	for key :=range listofKey{
		fmt.Println(listofKey[key],dictionary[listofKey[key]])
	}


}
