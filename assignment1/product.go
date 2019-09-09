package main

import "fmt"

func main() {

	a := integers("3141592653589793238462643383279502884197169399375105820974944592")
	b := integers("2718281828459045235360287471352662497757247093699959574966967627")
	//fmt.Println(1234*5678)

	fmt.Println(a)
	fmt.Println(b)

	sum := make([]int, len(a)+len(b))

	lenB := len(b)

	for i := 1; i <= lenB; i++ {

		v := product(a, b[lenB-i], i)
		add(sum, v)
	}

	for _, v := range sum {
		fmt.Print(v)
	}
}

func product(a []int, b, index int) []int {

	lenA := len(a)
	result := make([]int, lenA+index)

	for i := lenA - 1; i >= 0; i-- {
		v := a[i]*b + result[i+1]

		result[i] = v / 10
		result[i+1] = (v % 10)
	}

	return result
}

func add(sum, a []int) {

	for i, j := len(a)-1, len(sum)-1; i >= 0; {
		v := sum[j] + a[i]

		if v > 9 {
			sum[j-1]++
			sum[j] = v % 10
		} else {
			sum[j] = v
		}
		j--
		i--
	}
}

func integers(s string) []int {

	lenS := len(s)
	result := make([]int, lenS)

	for i, v := range s {

		result[i] = int(v - '0')
	}

	return result
}
