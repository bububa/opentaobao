package tbtrade

import (
	"sync"
)

// LogisticServiceTag 结构体
type LogisticServiceTag struct {
	// 物流服务下的标签属性,多个标签之间有&#34;;&#34;分隔
	ServiceTag string `json:"service_tag,omitempty" xml:"service_tag,omitempty"`
	// 消费者选快递请直接判断service_tag是否包含companyCode。而不要判断service_type
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
}

var poolLogisticServiceTag = sync.Pool{
	New: func() any {
		return new(LogisticServiceTag)
	},
}

// GetLogisticServiceTag() 从对象池中获取LogisticServiceTag
func GetLogisticServiceTag() *LogisticServiceTag {
	return poolLogisticServiceTag.Get().(*LogisticServiceTag)
}

// ReleaseLogisticServiceTag 释放LogisticServiceTag
func ReleaseLogisticServiceTag(v *LogisticServiceTag) {
	v.ServiceTag = ""
	v.ServiceType = ""
	poolLogisticServiceTag.Put(v)
}
