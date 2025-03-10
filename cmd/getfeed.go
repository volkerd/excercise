package cmd

import (
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"regexp"
)

type NewsList struct {
	News    []News `json:"news,omitempty"`
	BaseUrl string `json:"base_url,omitempty"`
}

type News struct {
	Published   string `json:"published,omitempty"`
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func getId(guid string) (string, error) {
	r, err := regexp.Compile("http://heise.de/-([[:digit:]]*)")
	if err != nil {
		log.Fatal(err)
	}
	m := r.FindStringSubmatch(guid)
	if len(m) != 2 {
		return "", errors.New("guid does not match")
	}
	return m[1], nil
}

func newsList(count int) NewsList {
	list := NewsList{BaseUrl: "http://heise.de/"}
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, item := range feed.Items {
		id, err := getId(item.GUID)
		if err == nil {
			n := News{item.Published, id, item.Title, item.Description}
			list.News = append(list.News, n)
		}
		if count > 0 && len(list.News) >= count {
			break
		}
	}
	return list
}

func newsSingle(id string) (News, error) {
	list := newsList(0)
	for _, item := range list.News {
		if item.Id == id {
			return item, nil
		}
	}
	return News{}, errors.New("id not found")
}
