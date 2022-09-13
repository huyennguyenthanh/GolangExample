package main

import "fmt"

func main() {

	// Khai báo một mảng integer có 5 phần tử.
	array_demo := [5]int{1, 2, 3, 4, 5}
	array_demo[2] = 6

	fmt.Printf("%v\n", array_demo) // [1 2 6 4 5]
	fmt.Print(array_demo)          // [1 2 6 4 5]

	// Khai báo một mảng con trỏ integer có 5 phần tử.
	// Khởi tạo giá trị cho phần tử có index là 0 và 1.
	array_pointer := [2]*int{0: new(int), 1: new(int)}
	*array_pointer[0] = 10

	fmt.Printf("%v\n", array_pointer) // [0xc0000180c0 0xc0000180c8]
	fmt.Println(array_pointer)        // [0xc0000180c0 0xc0000180c8]

	for _, p := range array_pointer {
		fmt.Println(p, *p)
	}
	// 0xc0000180c0 10
	// 0xc0000180c8 0

	var array [3]string
	array[1] = "Hi"
	fmt.Print(array) // [ Hi ]

}
