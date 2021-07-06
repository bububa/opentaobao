package mozi

// GetEmployeeTenantAdminInfoResult 结构体
type GetEmployeeTenantAdminInfoResult struct {
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// responseMessage
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// responseCode
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否是管理员
	Admin bool `json:"admin,omitempty" xml:"admin,omitempty"`
	// 是否主管理员
	Primary bool `json:"primary,omitempty" xml:"primary,omitempty"`
}
