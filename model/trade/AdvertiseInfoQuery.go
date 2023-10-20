package trade

import (
	"sync"
)

// AdvertiseInfoQuery 结构体
type AdvertiseInfoQuery struct {
	// 用户id
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolAdvertiseInfoQuery = sync.Pool{
	New: func() any {
		return new(AdvertiseInfoQuery)
	},
}

// GetAdvertiseInfoQuery() 从对象池中获取AdvertiseInfoQuery
func GetAdvertiseInfoQuery() *AdvertiseInfoQuery {
	return poolAdvertiseInfoQuery.Get().(*AdvertiseInfoQuery)
}

// ReleaseAdvertiseInfoQuery 释放AdvertiseInfoQuery
func ReleaseAdvertiseInfoQuery(v *AdvertiseInfoQuery) {
	v.OpenId = ""
	v.OrderId = 0
	poolAdvertiseInfoQuery.Put(v)
}
