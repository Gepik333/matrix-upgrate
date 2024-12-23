package main

import "fmt"

// func for output
func output(slice []int, stolb int) {
	for index, el := range slice {
		if index%stolb == 0 {
			fmt.Println(" ")
			fmt.Print("", el, "")
		} else {
			fmt.Print(" ", el, "")
		}
	}
}

func main() {
	var stolb, strok int
	stolb = 4 // pillars
	strok = 4 // lines
	// хуй
	a := stolb * strok
	elements(a, stolb)

}

// func matrix addition

func pl_matrix(slice_a []int, slice_b []int, stolb int) {
	slice_c := make([]int, len(slice_a))
	for i := 0; i < len(slice_a); i++ {
		slice_c[i] = slice_a[i] + slice_b[i]
	}
	output(slice_c, stolb)
}

// function for entering elements
func elements(a int, stolb int) {
	slice_a := make([]int, a)
	slice_b := make([]int, a)
	for i := 0; i < a; i++ { // append element for slice_a
		var el_a int
		fmt.Print("enter el #", i+1, ": ")
		fmt.Scan(&el_a)
		slice_a[i] = el_a
		fmt.Println(" ")
	}
	for i := 0; i < a; i++ {
		var el_b int
		fmt.Print("enter el #", i+1, ": ")
		fmt.Scan(&el_b)
		slice_b[i] = el_b
		fmt.Println(" ")
	}
	pl_matrix(slice_a, slice_b, stolb)
}
