package drugtrace

// AlibabaAlihealthDrugDownloadEntlistResult 
type AlibabaAlihealthDrugDownloadEntlistResult struct {
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // model
    Model   *DataEntTaskDto `json:"model,omitempty" xml:"model,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
}
