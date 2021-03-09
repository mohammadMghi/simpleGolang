package entity

type Post struct {
	ID int `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title,omitempty" bson:"title,omitempty" `
	Text string `json:"text" bson:"text,omitempty"`
}

