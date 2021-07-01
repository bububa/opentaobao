package moziacl

// GrantRolesResult 结构体
type GrantRolesResult struct {
	// 是否调用成功，成功则为true，否则为false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息，若失败，则返回失败信息
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 扩展字段，与入参扩展字段值相同
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}
