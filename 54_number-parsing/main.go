package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	fmt.Println(reflect.TypeOf(u))

	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	fmt.Println(reflect.TypeOf(k))

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e))
}
