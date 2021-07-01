package moziacl

// RevokeRolesResult 结构体
type RevokeRolesResult struct {
	// 是否调用成功，成功则为true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 请求唯一id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应message，若失败则返回失败原因
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 扩展字段，与入参扩展字段值相同
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}
