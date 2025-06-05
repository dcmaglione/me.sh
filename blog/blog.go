package blog

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}

func NewPost(content string) Post {
	return Post{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Content:   content,
	}
}

func Serialize(post Post) ([]byte, error) {
	return json.Marshal(post)
}

func Deserialize(data []byte) (Post, error) {
	var p Post
	err := json.Unmarshal(data, &p)
	return p, err
}
