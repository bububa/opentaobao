package alsc

// AlibabaAlscSaasCodecCodeAttrsQueryResult 结构体
type AlibabaAlscSaasCodecCodeAttrsQueryResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回素材id
	Data *CodeBizAttributeDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
