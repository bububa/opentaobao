package taotv

// TaobaotaotvvideoplaylistgetResult 结构体
type TaobaotaotvvideoplaylistgetResult struct {
	// 数据列表
	ModelList []TaobaotaotvvideoplaylistgetModel `json:"model_list,omitempty" xml:"model_list>taobaotaotvvideoplaylistget_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
