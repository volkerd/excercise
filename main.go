package main

import "fmt"
import "github.com/mmcdole/gofeed"

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	fmt.Printf("feed %T\n", feed.Items)
	for i := 0; i < 50; i++ {
		fmt.Printf("%v: %s\n", i+1, feed.Items[i].Title)
	}

}
