package promotion

import (
	"sync"
)

// TmallPromotionCouponQueryResult 结构体
type TmallPromotionCouponQueryResult struct {
	// data
	DataList []TmallPromotionCouponQueryData `json:"data_list,omitempty" xml:"data_list>tmall_promotion_coupon_query_data,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTmallPromotionCouponQueryResult = sync.Pool{
	New: func() any {
		return new(TmallPromotionCouponQueryResult)
	},
}

// GetTmallPromotionCouponQueryResult() 从对象池中获取TmallPromotionCouponQueryResult
func GetTmallPromotionCouponQueryResult() *TmallPromotionCouponQueryResult {
	return poolTmallPromotionCouponQueryResult.Get().(*TmallPromotionCouponQueryResult)
}

// ReleaseTmallPromotionCouponQueryResult 释放TmallPromotionCouponQueryResult
func ReleaseTmallPromotionCouponQueryResult(v *TmallPromotionCouponQueryResult) {
	v.DataList = v.DataList[:0]
	v.ErrorMsg = ""
	v.ResultCode = ""
	poolTmallPromotionCouponQueryResult.Put(v)
}
