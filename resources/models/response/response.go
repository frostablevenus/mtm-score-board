package response

type Record struct {
	Name  string `json:"name" binding:"required"`
	Score int    `json:"score" binding:"required"`
}

type Player struct {
	Name  string `json:"name" binding:"required"`
	Score []int  `json:"score" binding:"required"`
}
