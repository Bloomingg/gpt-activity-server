package request

type CreateActivityReq struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Total       int    `json:"total" form:"total" binding:"required"`
	Current     int    `json:"current" form:"current" gorm:"default:0;not null" `
	StartTime   uint   `json:"start_time" form:"start_time" binding:"required" `
	EndTime     uint   `json:"end_time" form:"end_time" binding:"required" `
	Location    string `json:"location" form:"location" binding:"required" `
	Latng       string `json:"latng" form:"latng"`
	Description string `json:"description" form:"description"`
	Banner      string `json:"banner" form:"banner"`
	Poster      string `json:"poster" form:"poster"`
}

type UpdateActivityReq struct {
	CreateActivityReq
	ID uint `json:"id" form:"id" binding:"required"`
}

type GetActivityReq struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
