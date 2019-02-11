package hackernews

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	baseURI        = "https://hacker-news.firebaseio.com"
	apiVersion     = "v0"
	defaultTimeout = 15 * time.Second
)

// New creates a new hackernews client using the given http client
// If no http client is provided, will use a plain http client with a
// 15 second timeout
func New(c *http.Client) Hackernews {
	if c != nil {
		return &hackernews{
			client: c,
		}
	}
	return &hackernews{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// Hackernews is an interface for interacting with the hackernews public API
type Hackernews interface {
	GetItem(id int) (*Item, error)
	GetUser(id string) (*User, error)
	MaxItemID() (int, error)
	TopStories() ([]int, error)
	NewStories() ([]int, error)
	AskStories() ([]int, error)
	ShowStories() ([]int, error)
	JobStories() ([]int, error)
}
type hackernews struct {
	client *http.Client
}

func (hn *hackernews) GetItem(id int) (*Item, error) {
	var story Item
	url := fmt.Sprintf("%s/%s/item/%d.json", baseURI, apiVersion, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&story); err != nil {
		return nil, err
	}
	return &story, nil
}

// GetUser will get user by id
func (hn *hackernews) GetUser(id string) (*User, error) {
	var user User
	url := fmt.Sprintf("%s/%s/user/%s.json", baseURI, apiVersion, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil

}

func (hn *hackernews) MaxItemID() (int, error) {
	res, err := hn.client.Get(fmt.Sprintf("%s/%s/maxitem.json", baseURI, apiVersion))
	if err != nil {
		return 0, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		resBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return 0, err
		}
		maxItemCount, err := strconv.Atoi(string(resBytes))
		if err != nil {
			return 0, err
		}
		return maxItemCount, nil
	default:
		return 0, errors.New("client returned code " + string(res.StatusCode) + ":" + res.Status)
	}
}
func (hn *hackernews) TopStories() ([]int, error) {
	url := fmt.Sprintf("%s/%v/topstories.json", baseURI, apiVersion)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var storyIDs []int
	if err := json.NewDecoder(res.Body).Decode(&storyIDs); err != nil {
		return nil, err
	}
	return storyIDs, nil
}
func (hn *hackernews) NewStories() ([]int, error) {
	url := fmt.Sprintf("%s/%v/newstories.json", baseURI, apiVersion)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var storyIDs []int
	if err := json.NewDecoder(res.Body).Decode(&storyIDs); err != nil {
		return nil, err
	}
	return storyIDs, nil
}
func (hn *hackernews) AskStories() ([]int, error) {
	url := fmt.Sprintf("%s/%v/askstories.json", baseURI, apiVersion)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var storyIDs []int
	if err := json.NewDecoder(res.Body).Decode(&storyIDs); err != nil {
		return nil, err
	}
	return storyIDs, nil
}
func (hn *hackernews) ShowStories() ([]int, error) {
	url := fmt.Sprintf("%s/%v/showstories.json", baseURI, apiVersion)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var storyIDs []int
	if err := json.NewDecoder(res.Body).Decode(&storyIDs); err != nil {
		return nil, err
	}
	return storyIDs, nil
}
func (hn *hackernews) JobStories() ([]int, error) {
	url := fmt.Sprintf("%s/%v/jobstories.json", baseURI, apiVersion)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var storyIDs []int
	if err := json.NewDecoder(res.Body).Decode(&storyIDs); err != nil {
		return nil, err
	}
	return storyIDs, nil
}
