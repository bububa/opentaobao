package mozi

// GetTenantByIdResult 结构体
type GetTenantByIdResult struct {
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// responseMessage
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// responseCode
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// data数据
	Data *Tenant `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
