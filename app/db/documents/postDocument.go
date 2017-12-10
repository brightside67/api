package documents

//PostDocument dad
type PostDocument struct {
	ID      string `bson:"_id,omitempty"`
	Title   string
	Content string
	Date    interface{}
}
