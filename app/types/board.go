package types

type AddReq struct {
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	Img      string `json:"img" form:"img"`
	UserName string `json:"userName" form:"userName"`
}
