package cmd

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/volkerd/exercise/common"
	"os"
)

func describe(id string) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, item := range feed.Items {
		if item.GUID == id {
			fmt.Println(item.Title, " @ ", common.FormatDate(item.Published))
			fmt.Println(" ")
			fmt.Println(item.Description)
			return
		}
	}
	fmt.Println(id + " not found.")
}
