package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotionCouponUse 券核销接口
// tmall.promotion.coupon.use
//
// 核销用户的一张优惠券，返回核销结果
func TmallPromotionCouponUse(clt *core.SDKClient, req *promotion.TmallPromotionCouponUseAPIRequest, resp *promotion.TmallPromotionCouponUseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
