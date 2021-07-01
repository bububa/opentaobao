package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponSpreadApplyAPIRequest
普通发券 API请求
alibaba.wdk.coupon.spread.apply

优惠券发放 */
type AlibabaWdkCouponSpreadApplyAPIRequest struct {
	model.Params
	// 参数对象
	_paramWdkCouponApplyParam *WdkCouponApplyParam
}

// New
