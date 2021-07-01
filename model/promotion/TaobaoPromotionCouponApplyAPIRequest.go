package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionCouponApplyAPIRequest
优惠券领取 API请求
taobao.promotion.coupon.apply

优惠券领取 */
type TaobaoPromotionCouponApplyAPIRequest struct {
	model.Params
	// 卖家id
	_sellerId string
	// 传播id
	_spreadId string
}

// New
