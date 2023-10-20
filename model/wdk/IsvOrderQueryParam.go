package wdk

import (
	"sync"
)

// IsvOrderQueryParam 结构体
type IsvOrderQueryParam struct {
	// 商家外部门店编码
	OutShopCode string `json:"out_shop_code,omitempty" xml:"out_shop_code,omitempty"`
	// 外部订单id
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

var poolIsvOrderQueryParam = sync.Pool{
	New: func() any {
		return new(IsvOrderQueryParam)
	},
}

// GetIsvOrderQueryParam() 从对象池中获取IsvOrderQueryParam
func GetIsvOrderQueryParam() *IsvOrderQueryParam {
	return poolIsvOrderQueryParam.Get().(*IsvOrderQueryParam)
}

// ReleaseIsvOrderQueryParam 释放IsvOrderQueryParam
func ReleaseIsvOrderQueryParam(v *IsvOrderQueryParam) {
	v.OutShopCode = ""
	v.OutOrderId = ""
	poolIsvOrderQueryParam.Put(v)
}
