package tmallnr

import (
	"sync"
)

// TmallNrtCouponInstanceQueryResult 结构体
type TmallNrtCouponInstanceQueryResult struct {
	// 券实例实体
	Models []string `json:"models,omitempty" xml:"models>string,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTmallNrtCouponInstanceQueryResult = sync.Pool{
	New: func() any {
		return new(TmallNrtCouponInstanceQueryResult)
	},
}

// GetTmallNrtCouponInstanceQueryResult() 从对象池中获取TmallNrtCouponInstanceQueryResult
func GetTmallNrtCouponInstanceQueryResult() *TmallNrtCouponInstanceQueryResult {
	return poolTmallNrtCouponInstanceQueryResult.Get().(*TmallNrtCouponInstanceQueryResult)
}

// ReleaseTmallNrtCouponInstanceQueryResult 释放TmallNrtCouponInstanceQueryResult
func ReleaseTmallNrtCouponInstanceQueryResult(v *TmallNrtCouponInstanceQueryResult) {
	v.Models = v.Models[:0]
	v.Code = ""
	v.Message = ""
	poolTmallNrtCouponInstanceQueryResult.Put(v)
}
