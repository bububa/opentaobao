package qimen

import (
	"sync"
)

// TaobaoQimenWarehouseinfoQueryRequest 结构体
type TaobaoQimenWarehouseinfoQueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenWarehouseinfoQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolTaobaoQimenWarehouseinfoQueryRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWarehouseinfoQueryRequest)
	},
}

// GetTaobaoQimenWarehouseinfoQueryRequest() 从对象池中获取TaobaoQimenWarehouseinfoQueryRequest
func GetTaobaoQimenWarehouseinfoQueryRequest() *TaobaoQimenWarehouseinfoQueryRequest {
	return poolTaobaoQimenWarehouseinfoQueryRequest.Get().(*TaobaoQimenWarehouseinfoQueryRequest)
}

// ReleaseTaobaoQimenWarehouseinfoQueryRequest 释放TaobaoQimenWarehouseinfoQueryRequest
func ReleaseTaobaoQimenWarehouseinfoQueryRequest(v *TaobaoQimenWarehouseinfoQueryRequest) {
	v.OwnerCode = ""
	v.ExtendProps = nil
	poolTaobaoQimenWarehouseinfoQueryRequest.Put(v)
}
