package cmd

import (
	"fmt"
)

func describe(id string) {
	item, err := newsSingle(id)
	if err != nil {
		fmt.Println(id + " not found.")
		return
	}
	fmt.Println(item.Title, " @ ", formatDate(item.Published))
	fmt.Println(" ")
	fmt.Println(item.Description)
}
