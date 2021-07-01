package eleenterprisecoupon

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseCouponGetAPIRequest
获取用户优惠券 API请求
alibaba.ele.enterprise.coupon.get

获取用户优惠券 */
type AlibabaEleEnterpriseCouponGetAPIRequest struct {
	model.Params
	// 手机号
	_phone string
}

// New
