package wdk

import (
	"sync"
)

// IsvQueryCouponParam 结构体
type IsvQueryCouponParam struct {
	// umpId列表，最多支持一次批量查询20个
	UmpIdList []int64 `json:"ump_id_list,omitempty" xml:"ump_id_list>int64,omitempty"`
}

var poolIsvQueryCouponParam = sync.Pool{
	New: func() any {
		return new(IsvQueryCouponParam)
	},
}

// GetIsvQueryCouponParam() 从对象池中获取IsvQueryCouponParam
func GetIsvQueryCouponParam() *IsvQueryCouponParam {
	return poolIsvQueryCouponParam.Get().(*IsvQueryCouponParam)
}

// ReleaseIsvQueryCouponParam 释放IsvQueryCouponParam
func ReleaseIsvQueryCouponParam(v *IsvQueryCouponParam) {
	v.UmpIdList = v.UmpIdList[:0]
	poolIsvQueryCouponParam.Put(v)
}
