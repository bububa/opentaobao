package icbuseller

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerCouponAuthVerify 优惠券校验
// alibaba.seller.coupon.auth.verify
//
// 优惠券校验
func AlibabaSellerCouponAuthVerify(ctx context.Context, clt *core.SDKClient, req *icbuseller.AlibabaSellerCouponAuthVerifyAPIRequest, resp *icbuseller.AlibabaSellerCouponAuthVerifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
