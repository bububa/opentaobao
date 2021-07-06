package promotion

// BenefitTemplateVo 结构体
type BenefitTemplateVo struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 外部关联id（如支付宝红包模板id）
	OutObjectId string `json:"out_object_id,omitempty" xml:"out_object_id,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 外部关联名称
	OutObjectName string `json:"out_object_name,omitempty" xml:"out_object_name,omitempty"`
	// 扩展信息
	ExtendFeature string `json:"extend_feature,omitempty" xml:"extend_feature,omitempty"`
	// 关联活动数
	RelateActivityNum int64 `json:"relate_activity_num,omitempty" xml:"relate_activity_num,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 权益id
	BenefitId int64 `json:"benefit_id,omitempty" xml:"benefit_id,omitempty"`
	// 可用总数
	ValidTotalNum int64 `json:"valid_total_num,omitempty" xml:"valid_total_num,omitempty"`
	// 面额
	Denomination int64 `json:"denomination,omitempty" xml:"denomination,omitempty"`
}
