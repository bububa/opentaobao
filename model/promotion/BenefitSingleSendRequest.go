package promotion

// BenefitSingleSendRequest 结构体
type BenefitSingleSendRequest struct {
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 幂等ID，代表一次发放。如，超时重试，需要保证幂等ID不变。不同的幂等ID代表两次不同得发放，因业务方更换幂等ID导致的超发问题由业务方负责
	UniqueId string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
	// 活动详情id
	DetailId int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// 活动id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 权益发放数量
	SendCount int64 `json:"send_count,omitempty" xml:"send_count,omitempty"`
}
