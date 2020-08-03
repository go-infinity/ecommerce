package main

import "fmt"

func main() {
	//Slice Example
	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for i, val := range s {
	// 	fmt.Printf("The value at %d index is %d\n", i, val)
	// }
	// fmt.Println("-----------------------")
	// //Slice Example with conditions
	// s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("The value at %d index is %d\n", i, s1[i])
	// }
	// fmt.Println("--------------------")

	// //Map Example
	// m := map[string]string{"A": "Apple", "B": "Ball", "C": "Cat"}
	// for key, val := range m {
	// 	fmt.Printf("The value of Key %s is %s\n", key, val)
	// }
	printstars(9)

}

func printstars(height int) {
	for i := 1; i <= height; i++ {
		k := 0

		for j := 1; j <= 2*i; j++ {
			//
			if k <= height/2+j+i && k > height/2-j-i {
				fmt.Printf("*")
			} else {
				fmt.Print(" ")
			}
			// for k := 0; k != 2*i-1; {
			// 	fmt.Printf("*")
			// }
		}
		fmt.Println(i)

	}
}
