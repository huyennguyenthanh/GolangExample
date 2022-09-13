package main

import pretty "github.com/inancgumus/prettyslice"

func main() {
	nums := []int{1, 9, 5, 6, 4, 8}
	odds := nums[:3]
	evens := nums[3:]

	nums[1], nums[3] = 9, 6
	pretty.Show("nums", nums)
	pretty.Show("odds : nums[:3]", odds)
	pretty.Show("evens: nums[3:]", evens)
}

// To run this file, following:
// go env -w GO111MODULE=auto
// git mod init prettyslice
// git mod tidy
// go run array.go

// source code: https://github.com/inancgumus/prettyslice
