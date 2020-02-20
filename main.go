package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func lendAndUpper(name string) (length int, upperCase string) {
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) (result int) {

	/* for - range */
	// for _, number := range numbers {
	// 	result += number
	// }

	/* for loop */
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return
}

func canIDrink(age int) bool {

	/* if - else */
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true

	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 19:
		return true
	case 50:
		return false
	}
	return false
}

func main() {
	// totalLength, upperCase := lendAndUpper("paul")
	// fmt.Println(totalLength, upperCase)
	// repeatMe("abc", "def", "ghi", "jkl")

	// totalNum := superAdd(1, 2, 3, 4, 5)
	// fmt.Println(totalNum)

	// fmt.Println(canIDrink(18))

	// a := 2
	// b := &a
	// a = 10
	// *b = 20
	// fmt.Println(a, *b)

	// names := []string{"abc", "def", "ghl"}
	// names = append(names, "ddd")
	// fmt.Println(names)

	/* map => map[key]value{key: value} */
	// paul := map[string]string{"name": "paul", "age": "10"}
	// for key, _ := range paul {
	// 	fmt.Println(key)
	// }

	/* struct */
	favFood := []string{"kimchi", "gogi"}
	paul := person{name: "paul", age: 18, favFood: favFood}
	fmt.Println(paul.name)
}
