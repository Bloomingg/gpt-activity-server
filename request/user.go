package request

type UserReq struct {
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
	Work     string `json:"work" form:"work"`
	Phone    string `json:"phone" form:"phone"`
	Activity uint   `json:"activity" form:"activity" binding:"required"`
}
