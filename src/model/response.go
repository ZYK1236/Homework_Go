package model

type ResponseModel struct {
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	Code       int         `json:"code"`
	TotalCount int         `json:"totalcount"`
}
