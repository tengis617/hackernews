# hackernews
A go client for [hackernews](https://github.com/HackerNews/API)


```go
package main

import (
	"fmt"
	"log"

	"github.com/tengis617/hackernews"
)

func main() {
	hn := hackernews.New(nil)
	storyIds, err := hn.TopStories()
	if err != nil {
		log.Fatal(err)
	}
	story, err := hn.GetItem(storyIds[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("todays top story:")
	fmt.Println(story.Title)
	fmt.Println(story.URL)
}
```