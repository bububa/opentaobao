package nlife

// DiscountMemo 结构体
type DiscountMemo struct {
	// 文案编号
	MemoId string `json:"memo_id,omitempty" xml:"memo_id,omitempty"`
	// 文案描述
	MemoDesc string `json:"memo_desc,omitempty" xml:"memo_desc,omitempty"`
}
