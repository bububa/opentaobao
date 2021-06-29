package tmallservice

// TmallServicecenterServicestoreUpdatestatusResult 
type TmallServicecenterServicestoreUpdatestatusResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    MsgCode   int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
