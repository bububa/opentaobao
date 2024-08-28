package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderCouponValidate 携带券的试单接口
// alitrip.merchant.galaxy.order.coupon.validate
//
// 试单时可以使用优惠券，返回一个原价，一个折扣价
func AlitripMerchantGalaxyOrderCouponValidate(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderCouponValidateAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderCouponValidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
