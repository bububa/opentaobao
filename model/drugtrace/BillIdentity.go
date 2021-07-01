package drugtrace

// BillIdentity 结构体
type BillIdentity struct {
	// 单据iD
	BillId int64 `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// 单据类型
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
}
