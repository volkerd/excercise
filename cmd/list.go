package cmd

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/volkerd/exercise/common"
	"os"
)

func list() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("feed %v\n", len(feed.Items))
	for i := 0; i < 5; i++ {
		item := feed.Items[i]
		fmt.Printf("%v: %s @ %s (%v)\n", i+1, item.Title, common.FormatDate(item.Published), item.GUID)
	}
}
