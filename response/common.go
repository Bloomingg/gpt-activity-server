package response

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

type DataList struct {
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
	Success bool        `json:"success"`
}

type GetDataListResponse struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   DataList `json:"data"`
	Error  string   `json:"error"`
}

type GetTokenData struct {
	UserID           string `json:"userID"`
	Token            string `json:"token"`
	ServerToken      string `json:"server_token"`
	ExpiredTime      int64  `json:"expiredTime"`
	MerchantManageID uint   `json:"merchant_manage_id"`
}

type IMApiCommonResponse struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

type GetTokenResponse struct {
	IMApiCommonResponse
	Data GetTokenData `json:"data"`
}
