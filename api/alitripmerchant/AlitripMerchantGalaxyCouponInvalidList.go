package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCouponInvalidList 用户失效优惠券列表
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
func AlitripMerchantGalaxyCouponInvalidList(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCouponInvalidListAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCouponInvalidListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
