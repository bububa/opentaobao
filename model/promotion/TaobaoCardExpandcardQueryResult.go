package promotion

// TaobaoCardExpandcardQueryResult 
type TaobaoCardExpandcardQueryResult struct {

    // 0为成功，其他为失败
    Code   int64 `json:"code,omitempty"`

    // 卡信息
    Models   []ExpandCardVo `json:"models,omitempty"`

    // debugInfo
    DebugInfo   string `json:"debug_info,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 错误级别
    ErrorLevel   string `json:"error_level,omitempty"`

}