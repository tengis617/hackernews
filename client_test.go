package hackernews

import (
	"testing"
)

func TestIntegration(t *testing.T) {
	hn := New(nil)

	maxItemCount, err := hn.MaxItemID()
	if err != nil {
		t.Fatal(err)
	}
	if maxItemCount == 0 {
		t.Fatalf("expecting max item id to not be 0")
	}

	topStories, err := hn.TopStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(topStories) == 0 {
		t.Fatalf("expecting top stories to not be 0")
	}

	i, err := hn.GetItem(topStories[0])
	if err != nil {
		t.Fatal(err)
	}
	if i == nil {
		t.Fatalf("item with id %d returned nil", topStories[0])
	}

	askStories, err := hn.AskStories()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("got %d stories from askStories", len(askStories))

	if len(askStories) == 0 {
		t.Fatalf("expecting top stories to not be 0")
	}

	jobStories, err := hn.JobStories()
	if err != nil {
		t.Fatal(err)
	}
	if len(jobStories) == 0 {
		t.Fatalf("expecting job stories to not be 0")
	}

	newStories, err := hn.NewStories()
	if err != nil {
		t.Fatal(err)
	}
	if len(newStories) == 0 {
		t.Fatalf("expecting new stories to not be 0")
	}

	showStories, err := hn.ShowStories()
	if err != nil {
		t.Fatal(err)
	}
	if len(showStories) == 0 {
		t.Fatalf("expecting show stories to not be 0")
	}

}
