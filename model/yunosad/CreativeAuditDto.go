package yunosad

// CreativeAuditDTO 
type CreativeAuditDTO struct {
    // 广告创意id
    CreativeId   string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
    // 是否审核通过，WAITING等待审核，PASS通过，REFUSE拒绝
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 创意级别
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    // 拒绝原因
    RefuseCause   string `json:"refuse_cause,omitempty" xml:"refuse_cause,omitempty"`
}
