package main

import (
	"fmt"
)

func arrays(n int) {
	// var name string
	var name1 string
	namelist := make([]string, n)

	// fmt.Printf("size of namelist 1 = %d ", len(namelist))

	for i := 0; i < n; i++ {
		fmt.Printf("Enter the name %d : ",i+1)
		fmt.Scanf("%s", &name1)
		namelist[i]=name1
		// namelist = append(namelist, name1)

	}

	// fmt.Printf("size of namelist 2 = %d ", len(namelist))

	// fmt.Print("Enter the name 2 : ")
	// fmt.Scanf("%s", &name1)
	// namelist = append(namelist, name1)

	fmt.Printf("size of namelist 3 = %d ", len(namelist))

	fmt.Println("Entered Names are :", namelist)

}
