package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoMobilePromotionCouponApply 优惠券领取(手淘专用)
// taobao.mobile.promotion.coupon.apply
//
// 优惠券领取
func TaobaoMobilePromotionCouponApply(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoMobilePromotionCouponApplyAPIRequest, resp *promotion.TaobaoMobilePromotionCouponApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
