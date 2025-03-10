package cmd

import (
	"fmt"
)

func list() {
	list := newsList(5)
	for i, item := range list.News {
		fmt.Printf("%v: %s @ %s (%v)\n", i+1, item.Title, formatDate(item.Published), item.Id)
	}
}
