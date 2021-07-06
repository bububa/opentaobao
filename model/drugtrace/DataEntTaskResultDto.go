package drugtrace

// DataEntTaskResultDto 结构体
type DataEntTaskResultDto struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
