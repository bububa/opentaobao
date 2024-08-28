package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCouponValidList 用户有效优惠券列表
// alitrip.merchant.galaxy.coupon.valid.list
//
// 雅高小程序用户有效优惠券列表
func AlitripMerchantGalaxyCouponValidList(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCouponValidListAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCouponValidListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
