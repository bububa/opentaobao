package mydata

// ProductEffect 结构体
type ProductEffect struct {
	// 产品效果
	Effects []EffectEntity `json:"effects,omitempty" xml:"effects>effect_entity,omitempty"`
	// 返回结果数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
