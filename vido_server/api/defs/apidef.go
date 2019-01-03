package defs

// reqeusts
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd string `json:"pwd"`
}

//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

// Data model
type VideoInfo struct{
	Id string
	AuthorId int
	Name string
	DisplayCtine string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string
	TTL int64
}