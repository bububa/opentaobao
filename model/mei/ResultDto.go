package mei

// ResultDto 结构体
type ResultDto struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// total
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// result
	Result *MemberAccountDto `json:"result,omitempty" xml:"result,omitempty"`
}
