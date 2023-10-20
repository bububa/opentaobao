package promotion

import (
	"sync"
)

// Range 结构体
type Range struct {
	// 营销范围参与者id。即MarketingRangeDO中的participateId。
	ParticipateId int64 `json:"participate_id,omitempty" xml:"participate_id,omitempty"`
	// 营销范围参与者类型。MarketingRangeDO中的participateType。(1:商品;2:店铺;3:seller;4:sku;5:category;6:shopCategory)
	ParticipateType int64 `json:"participate_type,omitempty" xml:"participate_type,omitempty"`
}

var poolRange = sync.Pool{
	New: func() any {
		return new(Range)
	},
}

// GetRange() 从对象池中获取Range
func GetRange() *Range {
	return poolRange.Get().(*Range)
}

// ReleaseRange 释放Range
func ReleaseRange(v *Range) {
	v.ParticipateId = 0
	v.ParticipateType = 0
	poolRange.Put(v)
}
