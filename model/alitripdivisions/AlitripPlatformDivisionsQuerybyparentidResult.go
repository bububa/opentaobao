package alitripdivisions

// AlitripPlatformDivisionsQuerybyparentidResult 
type AlitripPlatformDivisionsQuerybyparentidResult struct {
    // http status code
    HttpStatusCode   int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
    // mapping code
    MappingCode   string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
    // model
    Model   *TrdiDivisionBasicListVo `json:"model,omitempty" xml:"model,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
