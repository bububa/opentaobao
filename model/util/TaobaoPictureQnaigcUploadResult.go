package util

// TaobaoPictureQnaigcUploadResult 结构体
type TaobaoPictureQnaigcUploadResult struct {
	// 上传额外信息
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// 上传结果状态信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 上传结果状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// http状态码
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 图片结果
	Model *FileDo `json:"model,omitempty" xml:"model,omitempty"`
}
