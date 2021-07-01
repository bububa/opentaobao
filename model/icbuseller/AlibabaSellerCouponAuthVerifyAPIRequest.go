package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSellerCouponAuthVerifyAPIRequest
优惠券校验 API请求
alibaba.seller.coupon.auth.verify

优惠券校验 */
type AlibabaSellerCouponAuthVerifyAPIRequest struct {
	model.Params
	// 服务代码
	_serviceCode string
	// 卡券验证码
	_couponSeqNumber string
}

// New
