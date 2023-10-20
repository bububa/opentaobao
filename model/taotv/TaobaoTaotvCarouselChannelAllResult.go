package taotv

// TaobaotaotvcarouselchannelallResult 结构体
type TaobaotaotvcarouselchannelallResult struct {
	// 频道列表
	ModelList []TaobaotaotvcarouselchannelallModel `json:"model_list,omitempty" xml:"model_list>taobaotaotvcarouselchannelall_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
