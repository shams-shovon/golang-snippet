package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

func main() {
	//-----------------------------------------
	// 1. Print Statements
	// Description: Demonstrates basic printing and formatting
	//-----------------------------------------
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s) // Output: Hello and welcome, gopher!
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 1000/i) // Output: i = 1000, 500, 333, 250, 200
	}
	fmt.Println("Sum:", 9+7)      // Output: Sum: 16
	fmt.Println("Divide:", 9/7.0) // Output: Divide: 1.2857142857142858

	//-----------------------------------------
	// 2. Variables and Constants
	// Description: Declares variables, constants, and prints them
	//-----------------------------------------
	var b, c = 3, 4
	fmt.Println(b + c) // Output: 7

	var d = true
	fmt.Println(d) // Output: true

	var e bool
	fmt.Println(e) // Output: false (default bool value)

	owner := "Shovon"
	fmt.Println(owner) // Output: Shovon

	const constVal = 3e20 / 50000
	fmt.Println(constVal)           // Output: 6e+15
	fmt.Println(int64(constVal))    // Output: 6000000000000000
	fmt.Println(math.Sin(constVal)) // Output: Sin value (float)

	//-----------------------------------------
	// 3. Loops and Conditions
	//-----------------------------------------
	for i := 0; i < 5; i++ {
		fmt.Println(i) // Output: 0 1 2 3 4
	}

	for j := range 6 {
		if j%2 != 0 {
			fmt.Println(j) // Output: 1 3 5
		}
	}

	num := 19
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has 2 digits") // Output: 19 has 2 digits
	}

	//-----------------------------------------
	// 4. Switch Statements and Type Checking
	//-----------------------------------------
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two") // Output: two
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Friday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println(time.Now()) // Output: current date & time

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool") // Output: bool
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Printf("%T\n", t)
		}
	}
	checkType(true)     // Output: bool
	checkType(1.0)      // Output: float64
	checkType("SHOVON") // Output: string

	//-----------------------------------------
	// 5. Arrays and Multi-dimensional Arrays
	//-----------------------------------------
	var arr [5]int
	fmt.Println(arr) // Output: [0 0 0 0 0]

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // Output: [1 2 3 4 5]

	bArr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(bArr) // Output: [1 2 3 4 5]

	cArr := [...]int{100, 3: 400, 4, 5}
	fmt.Println(len(cArr)) // Output: 5

	var matrix [2][3]int
	fmt.Println(matrix) // Output: [[0 0 0] [0 0 0]]

	f := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	f[0][0] = 100
	fmt.Println(f) // Output: [[100 2 3] [4 5 6]]

	g := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(g) // Output: [[1 2 3] [4 5 6]]

	//-----------------------------------------
	// 6. Slices
	//-----------------------------------------
	h := make([]string, 2, 5)
	h = append(h, "a", "b", "c", "d", "p", "q")
	h[len(h)-1] = "z"
	fmt.Println(len(h), cap(h)) // Output: 8 10
	fmt.Println(h)              // Output: ["" "" "a" "b" "c" "d" "p" "z"]

	l := h[2:]
	fmt.Println(l) // Output: ["a" "b" "c" "d" "p" "z"]

	sliceCopy := make([]string, len(l))
	copy(sliceCopy, l)
	fmt.Println("S:", sliceCopy) // Output: S: [a b c d p z]

	final := make([][]int, 3)
	for i := 0; i < 3; i++ {
		ln := i + 1
		final[i] = make([]int, ln)
		for j := 0; j < ln; j++ {
			final[i][j] = i + j
		}
	}
	fmt.Println(final) // Output: [[0] [1 2] [2 3 4]]

	check := make([][][]int, 5)
	for i := 0; i < 3; i++ {
		inn := i + 1
		check[i] = make([][]int, inn)
		for j := 0; j < inn; j++ {
			inner := inn + i + j
			check[i][j] = make([]int, inner)
			for k := 0; k < inner; k++ {
				check[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println(check) // Output: 3D nested structure with values

	//-----------------------------------------
	// 7. Maps
	//-----------------------------------------
	j := make(map[string]int)
	j["k1"] = 7
	j["k2"] = 13
	fmt.Println(j)                 // Output: map[k1:7 k2:13]
	fmt.Println(reflect.TypeOf(j)) // Output: map[string]int

	v := reflect.ValueOf(j).MapKeys()
	fmt.Println(v) // Output: [k1 k2]

	v1 := j["k1"]
	fmt.Println(v1) // Output: 7

	delete(j, "k1")
	fmt.Println(j) // Output: map[k2:13]

	j["k9"] = 19
	clear(j)
	fmt.Println(j) // Output: map[]

	p := map[string]int{"f1": 90, "f2": 7}
	p2 := map[string]int{"f1": 9, "f2": 7}
	if reflect.DeepEqual(p, p2) {
		fmt.Println("pass")
	}

	eMap := map[string]int{"f1": 90, "f2": 7}
	va, pr := eMap["f3"]                // not found
	val, pre := eMap["f1"]              // found
	fmt.Println(va, pr, "**", val, pre) // Output: 0 false ** 90 true
	if !pr {
		fmt.Println("pr") // Output: pr
	}

	//-----------------------------------------
	// 8. Functions and Variadic Functions
	//-----------------------------------------
	fmt.Println(summation(1, 2)) // Output: 3

	cSum, dSum := summation2(3, 4, 5)
	fmt.Println(cSum, dSum) // Output: 7 5

	fSum, mSum := sum(1, 5, 8, 4)
	fmt.Println(fSum, mSum) // Output: 18 4

	nums := []int{3, 6, 7, 3, 8}
	fmt.Println(sum(nums...)) // Output: 27 5

	m10 := make(map[int]int)
	m10[4] = 5
	m10[2] = 8
	fmt.Println(sum(m10[4])) // Output: 5 1

	//-----------------------------------------
	// 9. Closures
	//-----------------------------------------
}

// -----------------------------------------
// Function Definitions
// -----------------------------------------
func summation(a int, b int) int {
	return a + b
}

func summation1(a, b, c int) int {
	return a + b + c
}

func summation2(a, b, c int) (int, int) {
	return a + b, c
}

func sum(numbers ...int) (int, int) {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result, len(numbers)
}
