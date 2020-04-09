package main

import (
	"fmt"
	"os"

	"github.com/arlsclu7/golab/helper"
	"github.com/arlsclu7/golab/types"
)

func main() {
	helper.StartLine()
	// println(math.Pi)
	type Weekday int
	type Person struct {
		Name string
		Sex  int
	}
	var ccc = make(types.Complex)
	var ddd = make(Person)

	// *sp.Name = "Tom"
	fmt.Println(types.C1)
	os.Exit(1)
	const (
		Monday Weekday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	println(Tuesday)

	os.Exit(1)
	gg := []string{}
	// ff := []string{"1", "2", "3"}
	for i := 0; i < 6; i++ {
		gg = append(gg, "hello", "world")
		fmt.Println(i, gg, cap(gg), " ")
	}
	helper.DivideLine()
	dd := []int{}
	fmt.Println(cap(dd))
	for i := 0; i < 17; i++ {
		dd = append(dd, i)
		fmt.Println(i, dd, cap(dd), " ")
	}
	helper.DivideLine()
	ee := []bool{}
	fmt.Println(cap(ee))
	for i := 0; i < 5; i++ {
		ee = append(ee, true, false)
		fmt.Println(i, ee, cap(ee), " ")
	}
	var aa = "mystring"
	var bb = &aa
	println(aa, bb)
	var cc *int = new(int)
	*cc = 100
	println(*cc)

	helper.DivideLine()
	d := []int{}
	fmt.Println(cap(d))
	for i := 0; i < 16; i++ {
		d = append(d, i)
		fmt.Print(cap(d), " ")
	}
	fmt.Println(types.E1)
	helper.DivideLine()
	a := make([]int, 5, 6)
	b := a[0:4]
	a = append(a, 1)
	a[1] = 5
	fmt.Println(b)
	// [0 0 0 0]
	fmt.Println(a)
	// [0 5 0 0 0 1]</pre>
	// var a = types.Float0
	// var b byte = '\v'
	var c rune = '中'
	var mystr1 string = "hello中国"
	var mystr2 [5]byte = [5]byte{104, 101, 65, 108, 111}
	fmt.Println(types.Arr0)
	fmt.Println(types.Slice0[1:5])
	fmt.Println(types.Slice1)
	fmt.Printf("值为%d,类型为%T\n", types.Slice0[0:2], types.Slice0[0:2])
	println(len(mystr1))
	println(len(mystr2))
	helper.DivideLine()
	var x []int = append(types.Slice2, []int{4, 5, 6}...)
	var y []int = append(append([]int{-1, -2, -3}, types.Slice2...), types.Slice2[:3]...)
	var z []int = types.Slice0[0:2:5]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%s \n", mystr1)
	fmt.Printf("%s \n", mystr2)
	fmt.Printf("%.3f\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%U\n", c)
	fmt.Println(c)
	helper.NewLine()
}
