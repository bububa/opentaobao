package flight

// PenaltyDto 结构体
type PenaltyDto struct {
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 描述
	Descs string `json:"descs,omitempty" xml:"descs,omitempty"`
	// 截止时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 费用（分）
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// 航程下标
	OdIndex int64 `json:"od_index,omitempty" xml:"od_index,omitempty"`
	// 乘机人类型,1成人2儿童3婴儿
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 规则支持情况，0不支持，1结构化，2非结构化，3支持
	PenaltySupportType int64 `json:"penalty_support_type,omitempty" xml:"penalty_support_type,omitempty"`
	// 规则类型，0退票费用，1同舱改期费用，2误机罚金，3其他，4升舱费用，5已使用航段扣减金额，6退税费用，7签转，8特殊说明
	PenaltyType int64 `json:"penalty_type,omitempty" xml:"penalty_type,omitempty"`
	// 百分比
	Percent int64 `json:"percent,omitempty" xml:"percent,omitempty"`
	// 航段下标
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// 开始时间
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 航段状态2标准半程已使用,3部分已使用,4全程未使用
	TicketSegmentsStatus int64 `json:"ticket_segments_status,omitempty" xml:"ticket_segments_status,omitempty"`
	// 时间单位，0小时，1天，默认0小时
	TimeUnit int64 `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
}
