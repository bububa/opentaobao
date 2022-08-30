package topoaid

// OrderGroupResponse 结构体
type OrderGroupResponse struct {
	// 入参中的groupId
	GroupId string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 回传结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
