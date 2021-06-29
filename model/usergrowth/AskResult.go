package usergrowth

// AskResult 
type AskResult struct {
    // 智能出价信息
    SmartBid   string `json:"smart_bid,omitempty" xml:"smart_bid,omitempty"`
    // 追踪id
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // idfa的md5值， 32位小写
    IdfaMd5   string `json:"idfa_md5,omitempty" xml:"idfa_md5,omitempty"`
    // imei的md5值， 32位小写
    ImeiMd5   string `json:"imei_md5,omitempty" xml:"imei_md5,omitempty"`
    // 任务类型， 1： 拉新；2：促活
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
