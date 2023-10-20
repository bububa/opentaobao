package promotion

import (
	"sync"
)

// StrategyIssueResultDto 结构体
type StrategyIssueResultDto struct {
	// 展示面额单位
	DisplayAmountUnit string `json:"display_amount_unit,omitempty" xml:"display_amount_unit,omitempty"`
	// 扩展字段
	ExtraData string `json:"extra_data,omitempty" xml:"extra_data,omitempty"`
	// 发放时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 生效开始时间
	EffectiveStart string `json:"effective_start,omitempty" xml:"effective_start,omitempty"`
	// 外部发放实例id
	OuterInstanceId string `json:"outer_instance_id,omitempty" xml:"outer_instance_id,omitempty"`
	// 权益code
	BenefitCode string `json:"benefit_code,omitempty" xml:"benefit_code,omitempty"`
	// 生效结束时间
	EffectiveEnd string `json:"effective_end,omitempty" xml:"effective_end,omitempty"`
	// 展示门槛
	DisplayStartFee string `json:"display_start_fee,omitempty" xml:"display_start_fee,omitempty"`
	// 面额单位
	AmountUnit string `json:"amount_unit,omitempty" xml:"amount_unit,omitempty"`
	// 追踪数据
	TrackingData string `json:"tracking_data,omitempty" xml:"tracking_data,omitempty"`
	// 展示面额
	DisplayAmount string `json:"display_amount,omitempty" xml:"display_amount,omitempty"`
	// 生效时间类型
	EffectiveTimeMode string `json:"effective_time_mode,omitempty" xml:"effective_time_mode,omitempty"`
	// 权益标题
	BenefitTitle string `json:"benefit_title,omitempty" xml:"benefit_title,omitempty"`
	// 权益类型名称
	BenefitTypeName string `json:"benefit_type_name,omitempty" xml:"benefit_type_name,omitempty"`
	// 相对生效时间单位
	IntervalTimeUnit string `json:"interval_time_unit,omitempty" xml:"interval_time_unit,omitempty"`
	// 素材
	Material string `json:"material,omitempty" xml:"material,omitempty"`
	// 生效结束时间戳
	EffectiveEndTimestamp int64 `json:"effective_end_timestamp,omitempty" xml:"effective_end_timestamp,omitempty"`
	// 中奖记录id
	RecordId int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 生效开始时间戳
	EffectiveStartTimestamp int64 `json:"effective_start_timestamp,omitempty" xml:"effective_start_timestamp,omitempty"`
	// 面额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 门槛
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 相对生效时间长度
	EffectiveInterval int64 `json:"effective_interval,omitempty" xml:"effective_interval,omitempty"`
}

var poolStrategyIssueResultDto = sync.Pool{
	New: func() any {
		return new(StrategyIssueResultDto)
	},
}

// GetStrategyIssueResultDto() 从对象池中获取StrategyIssueResultDto
func GetStrategyIssueResultDto() *StrategyIssueResultDto {
	return poolStrategyIssueResultDto.Get().(*StrategyIssueResultDto)
}

// ReleaseStrategyIssueResultDto 释放StrategyIssueResultDto
func ReleaseStrategyIssueResultDto(v *StrategyIssueResultDto) {
	v.DisplayAmountUnit = ""
	v.ExtraData = ""
	v.IssueTime = ""
	v.BenefitType = ""
	v.EffectiveStart = ""
	v.OuterInstanceId = ""
	v.BenefitCode = ""
	v.EffectiveEnd = ""
	v.DisplayStartFee = ""
	v.AmountUnit = ""
	v.TrackingData = ""
	v.DisplayAmount = ""
	v.EffectiveTimeMode = ""
	v.BenefitTitle = ""
	v.BenefitTypeName = ""
	v.IntervalTimeUnit = ""
	v.Material = ""
	v.EffectiveEndTimestamp = 0
	v.RecordId = 0
	v.EffectiveStartTimestamp = 0
	v.Amount = 0
	v.StartFee = 0
	v.EffectiveInterval = 0
	poolStrategyIssueResultDto.Put(v)
}
