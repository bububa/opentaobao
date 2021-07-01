package taotv

// TaobaoTaotvVideoPlaylistOttnavGetResult 结构体
type TaobaoTaotvVideoPlaylistOttnavGetResult struct {
	// model
	Model *PlayListNavRbo `json:"model,omitempty" xml:"model,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
