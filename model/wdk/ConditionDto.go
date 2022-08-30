package wdk

// ConditionDto 结构体
type ConditionDto struct {
	// 满元门槛值 -- 单位分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 件数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 是否满元门槛
	AmountAt bool `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
	// 是否满件门槛
	CountAt bool `json:"count_at,omitempty" xml:"count_at,omitempty"`
}
