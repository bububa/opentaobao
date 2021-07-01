package taotv

// TaobaoTaotvCarouselChannelAllResult 结构体
type TaobaoTaotvCarouselChannelAllResult struct {
	// 频道列表
	ModelList []TaobaoTaotvCarouselChannelAllModel `json:"model_list,omitempty" xml:"model_list>taobao_taotv_carousel_channel_all_model,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
