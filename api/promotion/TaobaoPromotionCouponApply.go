package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionCouponApply 优惠券领取
// taobao.promotion.coupon.apply
//
// 优惠券领取
func TaobaoPromotionCouponApply(clt *core.SDKClient, req *promotion.TaobaoPromotionCouponApplyAPIRequest, resp *promotion.TaobaoPromotionCouponApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
