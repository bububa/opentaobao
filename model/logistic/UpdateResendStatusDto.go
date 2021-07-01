package logistic

// UpdateResendStatusDto 
type UpdateResendStatusDto struct {
    // 补发单状态（-1=关闭，1=补发成功，2=部分成功）
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 主订单
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    // 描述
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // erp补发单
    ReissueId   string `json:"reissue_id,omitempty" xml:"reissue_id,omitempty"`
    // 平台补发单唯一标识
    SourceId   string `json:"source_id,omitempty" xml:"source_id,omitempty"`
}
