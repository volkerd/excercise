package cmd

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"os"
)

func describe(id string) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, element := range feed.Items {
		if element.GUID == id {
			fmt.Println(element.Title)
			fmt.Println(" ")
			fmt.Println(element.Description)
			return
		}
	}
	fmt.Println(id + " not found.")
}
