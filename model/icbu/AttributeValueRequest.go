package icbu

import (
	"sync"
)

// AttributeValueRequest 结构体
type AttributeValueRequest struct {
	// 选填；需要过滤的属性值id
	AttributeValueId []int64 `json:"attribute_value_id,omitempty" xml:"attribute_value_id>int64,omitempty"`
	// 选填；需要过滤的属性
	AttributeId []int64 `json:"attribute_id,omitempty" xml:"attribute_id>int64,omitempty"`
	// 必填；要查询的属性值所属发布类目
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}

var poolAttributeValueRequest = sync.Pool{
	New: func() any {
		return new(AttributeValueRequest)
	},
}

// GetAttributeValueRequest() 从对象池中获取AttributeValueRequest
func GetAttributeValueRequest() *AttributeValueRequest {
	return poolAttributeValueRequest.Get().(*AttributeValueRequest)
}

// ReleaseAttributeValueRequest 释放AttributeValueRequest
func ReleaseAttributeValueRequest(v *AttributeValueRequest) {
	v.AttributeValueId = v.AttributeValueId[:0]
	v.AttributeId = v.AttributeId[:0]
	v.CatId = 0
	poolAttributeValueRequest.Put(v)
}
