package simba

import (
	"sync"
)

// ExtraAttributes 结构体
type ExtraAttributes struct {
	// 商品在淘宝的发布时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品的累积销量
	SalesCount int64 `json:"sales_count,omitempty" xml:"sales_count,omitempty"`
}

var poolExtraAttributes = sync.Pool{
	New: func() any {
		return new(ExtraAttributes)
	},
}

// GetExtraAttributes() 从对象池中获取ExtraAttributes
func GetExtraAttributes() *ExtraAttributes {
	return poolExtraAttributes.Get().(*ExtraAttributes)
}

// ReleaseExtraAttributes 释放ExtraAttributes
func ReleaseExtraAttributes(v *ExtraAttributes) {
	v.PublishTime = ""
	v.Quantity = 0
	v.SalesCount = 0
	poolExtraAttributes.Put(v)
}
