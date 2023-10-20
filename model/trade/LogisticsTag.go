package trade

import (
	"sync"
)

// LogisticsTag 结构体
type LogisticsTag struct {
	// 服务标签
	LogisticServiceTagList []LogisticServiceTag `json:"logistic_service_tag_list,omitempty" xml:"logistic_service_tag_list>logistic_service_tag,omitempty"`
	// 主订单或子订单的订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolLogisticsTag = sync.Pool{
	New: func() any {
		return new(LogisticsTag)
	},
}

// GetLogisticsTag() 从对象池中获取LogisticsTag
func GetLogisticsTag() *LogisticsTag {
	return poolLogisticsTag.Get().(*LogisticsTag)
}

// ReleaseLogisticsTag 释放LogisticsTag
func ReleaseLogisticsTag(v *LogisticsTag) {
	v.LogisticServiceTagList = v.LogisticServiceTagList[:0]
	v.OrderId = ""
	poolLogisticsTag.Put(v)
}
