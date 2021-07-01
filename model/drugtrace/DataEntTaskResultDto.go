package drugtrace

// DataEntTaskResultDto 
type DataEntTaskResultDto struct {
    // model
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
