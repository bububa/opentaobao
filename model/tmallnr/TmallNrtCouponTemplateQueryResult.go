package tmallnr

import (
	"sync"
)

// TmallNrtCouponTemplateQueryResult 结构体
type TmallNrtCouponTemplateQueryResult struct {
	// 券模版数据
	Model []NrtCouponTemplateDto `json:"model,omitempty" xml:"model>nrt_coupon_template_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTmallNrtCouponTemplateQueryResult = sync.Pool{
	New: func() any {
		return new(TmallNrtCouponTemplateQueryResult)
	},
}

// GetTmallNrtCouponTemplateQueryResult() 从对象池中获取TmallNrtCouponTemplateQueryResult
func GetTmallNrtCouponTemplateQueryResult() *TmallNrtCouponTemplateQueryResult {
	return poolTmallNrtCouponTemplateQueryResult.Get().(*TmallNrtCouponTemplateQueryResult)
}

// ReleaseTmallNrtCouponTemplateQueryResult 释放TmallNrtCouponTemplateQueryResult
func ReleaseTmallNrtCouponTemplateQueryResult(v *TmallNrtCouponTemplateQueryResult) {
	v.Model = v.Model[:0]
	v.Message = ""
	v.Code = 0
	poolTmallNrtCouponTemplateQueryResult.Put(v)
}
