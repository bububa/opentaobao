package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeCouponSubmitAPIRequest
提交专车优惠券活动 API请求
alibaba.alihouse.newhome.coupon.submit

提交专车优惠券活动 */
type AlibabaAlihouseNewhomeCouponSubmitAPIRequest struct {
	model.Params
	// 打车券活动对象
	_couponDto *MarketingCouponDto
}

// New
