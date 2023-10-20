package eleenterprisecoupon

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

// AlibabaEleEnterpriseCouponGet 获取用户优惠券
// alibaba.ele.enterprise.coupon.get
//
// 获取用户优惠券
func AlibabaEleEnterpriseCouponGet(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaEleEnterpriseCouponGetAPIRequest, resp *eleenterprisecoupon.AlibabaEleEnterpriseCouponGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
