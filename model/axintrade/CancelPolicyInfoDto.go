package axintrade

// CancelPolicyInfoDto 结构体
type CancelPolicyInfoDto struct {
	// 提前小时
	Hour int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// 规则对应的值，可能是百分比、数值等
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}
