package tanx

// CreativeDto 
type CreativeDto struct {

    // 创意ID
    CreativeId   string `json:"creative_id,omitempty"`

    // 创意审核的状态（通过PASS,拒绝REFUSE,未审核WAITING）
    Status   string `json:"status,omitempty"`

    // 创意通过的等级，1表示一级创意，99表示普通创意
    Level   int64 `json:"level,omitempty"`

    // 创意拒绝的原因
    RefuseCause   string `json:"refuse_cause,omitempty"`

    // 广告位属性
    AdboardData   string `json:"adboard_data,omitempty"`

    // 创意审核信息列表
    CreativePublisherAuditDtoList   []CreativePublisherAuditDto `json:"creative_publisher_audit_dto_list,omitempty"`

}
