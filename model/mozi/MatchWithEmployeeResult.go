package mozi

// MatchWithEmployeeResult 结构体
type MatchWithEmployeeResult struct {
	// 未匹配上的人员code
	UnmatchedEmployeeCodes []string `json:"unmatched_employee_codes,omitempty" xml:"unmatched_employee_codes>string,omitempty"`
	// 匹配上的人员code
	MatchedEmployeeCodes []string `json:"matched_employee_codes,omitempty" xml:"matched_employee_codes>string,omitempty"`
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// responseMessage
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// responseCode
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
