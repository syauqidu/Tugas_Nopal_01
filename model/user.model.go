package model

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Followers int       `json:"followers"`
	Following int       `json:"following"`
	Posted    []Posting `json:"posted"`
}
