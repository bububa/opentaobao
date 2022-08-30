package charity

// BillDto 结构体
type BillDto struct {
	// 账单编号
	BillId string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// 账期
	BillTime int64 `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
}
