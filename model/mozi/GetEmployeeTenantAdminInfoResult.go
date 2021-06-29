package mozi

// GetEmployeeTenantAdminInfoResult 
type GetEmployeeTenantAdminInfoResult struct {
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // requestId
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否是管理员
    Admin   bool `json:"admin,omitempty" xml:"admin,omitempty"`
    // responseMessage
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    // 是否主管理员
    Primary   bool `json:"primary,omitempty" xml:"primary,omitempty"`
    // responseCode
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}
