package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyReceiveCouponByActivity 按活动Id领取优惠券
// alitrip.merchant.galaxy.receive.coupon.by.activity
//
// 雅高小程序按活动Id领取优惠券
func AlitripMerchantGalaxyReceiveCouponByActivity(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyReceiveCouponByActivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
