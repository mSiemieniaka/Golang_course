package main

import "fmt"

func main() {
	list1 := [3]string{"Football", "Basketball", "Coding"} //<-lista do podunku 1

	list2 := list1[1:] //<-lista do podunku 2 drugi i trzeci element
	//list3 := list1[0:2] //<-lista do podunku 3 to ta pierwsza
	//list4 := list1[:2] // <-lista do podunku 3 to ta druga
	list3 := list1[1:3]

	dynamicList := []string{"GoDeveloper", "Devops"}
	dynamicList[1] = "Cos innego"
	dynamicList = append(dynamicList, "Macbook PRO")

	fmt.Println(list1)
	fmt.Println(list1[0:1]) //<-lista do podunku 2 standalone
	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(dynamicList)
	//fmt.Println(list4)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line. <- DONE
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list <- DONE
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end) <- DONE
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array. <- DONE
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals) <- DONE
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array <- DONE
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
