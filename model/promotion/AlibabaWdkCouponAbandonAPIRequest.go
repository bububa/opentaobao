package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponAbandonAPIRequest
废券 API请求
alibaba.wdk.coupon.abandon

优惠券废弃 */
type AlibabaWdkCouponAbandonAPIRequest struct {
	model.Params
	// 废券参数
	_paramWdkCouponAbandonParam *WdkCouponAbandonParam
}

// New
