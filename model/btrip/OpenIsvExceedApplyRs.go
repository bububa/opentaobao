package btrip

// OpenIsvExceedApplyRs 结构体
type OpenIsvExceedApplyRs struct {
	// 商旅企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 出差原因
	BtripCause string `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	// 超标原因
	ExceedReason string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// 原差旅标准
	OriginStandard string `json:"origin_standard,omitempty" xml:"origin_standard,omitempty"`
	// 第三方用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 审批单提交时间
	SubmitTime string `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	// 第三方企业id
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 第三方出差审批单号
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 商旅审批单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 审批单状态 0:审批中 1:已同意 2:已拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 超标类型，1:折扣 2,8,10:时间 3,9,11:折扣和时间
	ExceedType int64 `json:"exceed_type,omitempty" xml:"exceed_type,omitempty"`
	// 意向出行信息
	ApplyIntentionInfoDo *ApplyIntentionInfoDo `json:"apply_intention_info_do,omitempty" xml:"apply_intention_info_do,omitempty"`
}
