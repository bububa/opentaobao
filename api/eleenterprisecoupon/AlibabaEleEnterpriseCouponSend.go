package eleenterprisecoupon

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

// AlibabaEleEnterpriseCouponSend 发放优惠券
// alibaba.ele.enterprise.coupon.send
//
// 发放优惠券
func AlibabaEleEnterpriseCouponSend(ctx context.Context, clt *core.SDKClient, req *eleenterprisecoupon.AlibabaEleEnterpriseCouponSendAPIRequest, resp *eleenterprisecoupon.AlibabaEleEnterpriseCouponSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
