package eleenterprisecoupon

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

// AlibabaEleEnterpriseCouponSend 发放优惠券
// alibaba.ele.enterprise.coupon.send
//
// 发放优惠券
func AlibabaEleEnterpriseCouponSend(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaEleEnterpriseCouponSendAPIRequest, resp *eleenterprisecoupon.AlibabaEleEnterpriseCouponSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
