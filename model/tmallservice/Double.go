package tmallservice

// Double 结构体
type Double struct {
	// 费用名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 费用金额（分）
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
