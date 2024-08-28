package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityCouponList 用户领券中心列表
// alitrip.merchant.galaxy.activity.coupon.list
//
// 雅高小程序用户领券中心列表
func AlitripMerchantGalaxyActivityCouponList(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityCouponListAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityCouponListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
