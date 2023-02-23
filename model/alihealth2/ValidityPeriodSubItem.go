package alihealth2

// ValidityPeriodSubItem 结构体
type ValidityPeriodSubItem struct {
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 失效日期
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 生产批号
	ProduceNo string `json:"produce_no,omitempty" xml:"produce_no,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
