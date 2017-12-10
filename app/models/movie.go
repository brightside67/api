package models

//Post hz
type Post struct {
	ID      string      `json:"id"`
	Title   string      `json:"title"`
	Content string      `json:"content"`
	Date    interface{} `json:"date"`
}

//NewPost hz
func NewPost(id string, title string, content string, time interface{}) *Post {
	return &Post{id, title, content, time}
}
