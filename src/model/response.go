package model

type ResponseModel struct {
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	Code       int         `json:"code"`
	TotalCount int         `json:"totalcount"`
}

func GetModel(isSuccess bool, data interface{}, count int) *ResponseModel {
	if isSuccess {
		res := new(ResponseModel)
		res.Code = 1
		res.Msg = "success"
		res.Data = data
		res.TotalCount = count

		return res
	} else {
		res := new(ResponseModel)
		res.Code = 0
		res.Msg = "error"
		res.Data = nil

		return res
	}
}
