package axindata

// PolicyVo 结构体
type PolicyVo struct {
	// 扣除值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 往前推小时
	Hour int64 `json:"hour,omitempty" xml:"hour,omitempty"`
}
