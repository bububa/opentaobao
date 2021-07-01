package txcs

// BizExtAttr 结构体
type BizExtAttr struct {
	// 汇总编码
	SummaryId string `json:"summary_id,omitempty" xml:"summary_id,omitempty"`
	// 账单创建时间
	BillCreateTime string `json:"bill_create_time,omitempty" xml:"bill_create_time,omitempty"`
}
