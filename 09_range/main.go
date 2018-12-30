package main

import "fmt"

func main() {
	ids := []int{23,34,45,56,67,78,89,90}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using ids
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with map
	emails := map[string]string{"bob":"bob@foo.com", "vikesh":"vikesh@foo.com"}

	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}

	// If only key is required
	for k := range emails {
		fmt.Println("Names " + k)
	}
}