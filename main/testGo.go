package main

import (
	//"../testpackage" //when the package "testpackage" doesn't exit in the GOPATH
	"container/list"
	"fmt"
	"math/big"
	"sync"
	"testpackage" //when the package "testpackage" add in the GOPATH
	"time"
	"unicode/utf8"
)

var (
	vara = 1
	//varb int = 2
)

const (
	a   = 1
	str = "hello world 我"
)

type Info struct {
	str string
	mu  sync.Mutex
}

func main() {
	testpackage.TestPrint()
	//newAndMake()
	//slinceAndArray()
	//mapTest()
	//testPackage()
	//fmt.Println("===")
	//fmt.Println(f(), "===")
	//a, b, c := compute(1, 2)
	//fmt.Println(a, b, c)
}

func newAndMake() {
	p1 := new(int)
	fmt.Println(p1)  //(*int)(0xc42000e250)
	fmt.Println(*p1) //0

	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2)           //(*int)(0xc42000e278)
	fmt.Printf("p2 point to --> %#v \n ", *p2) //0
}

func updateMu(st string, info *Info) {
	info.mu.Lock()
	info.str = st
	info.mu.Unlock()
}

func testPackage() {
	list1 := list.New()
	list1.PushBack(101)
	list1.PushBack(102)
	list1.PushBack(103)
	fmt.Println("length of list1:", list1.Len())
	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	info := new(Info)
	for i := 0; i < 10; i++ {
		go updateMu("adf"+string(i), info)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(info.str)

	im := big.NewInt(123456789012)
	in := big.NewInt(10)
	ip := big.NewInt(1)
	ip.Add(im, in)
	fmt.Println(ip)
}

func mapTest() {
	var map1 map[string]int
	//map1 = map[string]int{"one": 1}
	map1 = make(map[string]int)
	map1["two"] = 2
	fmt.Println(map1["two"])

	var map11 = map[string]int{"one": 1, "two": 2}
	fmt.Println(map11)

	if val, ok := map11["three"]; ok {
		fmt.Println(val)
	}
	delete(map11, "one")
	fmt.Println(map11)
}

func split(str string, ind int) (str1, str2 string) {
	slice := []byte(str)
	str1, str2 = string(slice[:ind]), string(slice[ind:])
	return
}

func addUsedot(arr ...float64) (sum1 float64) {
	for _, val := range arr {
		sum1 += val
	}
	return
}

func addUseArray(arr [4]float64) (sum1 float64) {
	for _, val := range arr {
		sum1 += val
	}
	return
}

func sum(a *[4]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}

func addUseSlice(arr []float64) (sum1 float64) {
	//arr[0] = 0.0
	for _, val := range arr {
		sum1 += val
	}
	return
}

func slinceAndArray() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrLazy1 = []int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrKeyValue1 = []string{1: "Chris", 3: "Ron"}
	fmt.Println(len(arrAge), len(arrLazy), len(arrLazy1), len(arrKeyValue), len(arrKeyValue1))

	fmt.Println(fmt.Sprintf("%T", arrAge))
	fmt.Println(fmt.Sprintf("%T", arrLazy))
	fmt.Println(fmt.Sprintf("%T", arrLazy1))
	fmt.Println(fmt.Sprintf("%T", arrKeyValue))
	fmt.Println(fmt.Sprintf("%T", arrKeyValue1))

	var arr1 [5]int
	for i := 0; i < 5; i++ {
		arr1[i] = i + 1
	}

	//s := [...]int{1, 2, 3, 4, 5, 6}
	var slice1 = arr1[1:3]
	for ind, val := range slice1 {
		fmt.Println(ind, val)
	}

	arr := []int{1, 2, 3}
	fmt.Println(arr)
	for ind, val := range arr {
		fmt.Println(ind, val)
	}
	pri(arr...)

	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	for i := 0; i < len(items); i++ {
		items[i] *= 2
	}
	fmt.Println(items)
	s := make([]int, 5)
	for i := 0; i < 5; i++ {
		s[i] = i
	}
	fmt.Println(len(s), cap(s), s)
	s = s[2:4]
	fmt.Println(len(s), cap(s), s)
	s = s[2:3]
	fmt.Println(len(s), cap(s), s)

	var arr11 = [4]float64{1.1, 2.2, 3.3, 4.5}
	fmt.Println(addUseArray(arr11))
	fmt.Println(addUsedot(1, 2, 3))
	//fmt.Println(addUsedot(arr11...))
	fmt.Println(sum(&arr11))
	fmt.Println(addUseSlice(arr11[:]))

	var slice11 = []int{1, 2, 3, 4, 5}
	slice11 = slice11[3:4]
	fmt.Println(slice11)

	for ind, val := range str {
		fmt.Printf("%d:%c ", ind, val)
	}
	fmt.Println()
	fmt.Println("size of str:", utf8.RuneCountInString(str))

	str := "12345678"
	str1, str2 := split(str, 4)
	fmt.Println(str1, str2)
}

func pri(arr ...int) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func compute(a int, b int) (sum int, dif int, mul int) {
	sum, dif, mul = a+b, a-b, a*b
	//return a + b, a - b, a * b
	return
}

func f() (ret int) {
	fmt.Println(ret)
	defer func() {
		ret++
		fmt.Println(ret)
	}()
	return 3
}
