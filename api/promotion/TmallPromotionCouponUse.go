package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotionCouponUse 券核销接口
// tmall.promotion.coupon.use
//
// 核销用户的一张优惠券，返回核销结果
func TmallPromotionCouponUse(clt *core.SDKClient, req *promotion.TmallPromotionCouponUseAPIRequest, session string) (*promotion.TmallPromotionCouponUseAPIResponse, error) {
	var resp promotion.TmallPromotionCouponUseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
