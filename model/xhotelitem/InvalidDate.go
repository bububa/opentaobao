package xhotelitem

// InvalidDate 结构体
type InvalidDate struct {
	// 活动失效开始时间
	InvalidFrom string `json:"invalid_from,omitempty" xml:"invalid_from,omitempty"`
	// 活动失效结束时间
	InvalidTo string `json:"invalid_to,omitempty" xml:"invalid_to,omitempty"`
}
