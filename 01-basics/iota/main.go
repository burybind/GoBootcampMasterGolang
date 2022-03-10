package main

import "fmt"

func main() {
	const (
		monday = iota +1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	const (
		EST = -(5 + iota)
		_
		MST = -7
		PST = -8
	)

	fmt.Println(EST, MST, PST)
}
