package campus

// IdentifyAuthResultDto 
type IdentifyAuthResultDto struct {
    // 业务id
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    // 鉴权结果码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 鉴权结果消息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 鉴权结果消息英文
    ErrorMsgEn   string `json:"error_msg_en,omitempty" xml:"error_msg_en,omitempty"`
    // 业务流水id
    EventTraceId   string `json:"event_trace_id,omitempty" xml:"event_trace_id,omitempty"`
}
