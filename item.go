package hackernews

// Item is an item on hackernews
type Item struct {
	ID          int  `json:"id"`
	Deleted     bool `json:"deleted"`
	Type        ItemType
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Poll        int    `json:"poll"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Parts       []int  `json:"parts"`
	Descendants int    `json:"descendants"`
}

// ItemType is the type of item. One of "job", "story", "comment", "poll", or "pollopt".
type ItemType string

const (
	// ItemTypeJob is for Jobs
	ItemTypeJob ItemType = "job"
	// ItemTypeStory is for stories
	ItemTypeStory ItemType = "story"
	// ItemTypeComment is for comments
	ItemTypeComment ItemType = "comment"
	// ItemTypePoll is for polls
	ItemTypePoll ItemType = "poll"
	// ItemTypePollOpt is for poll opt
	ItemTypePollOpt ItemType = "pollopt"
)
