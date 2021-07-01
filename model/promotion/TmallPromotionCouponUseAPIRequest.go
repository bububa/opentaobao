package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotionCouponUseAPIRequest
券核销接口 API请求
tmall.promotion.coupon.use

核销用户的一张优惠券，返回核销结果 */
type TmallPromotionCouponUseAPIRequest struct {
	model.Params
	// 扩展字段
	_extra string
	// 业务类型
	_bizType string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerId string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerNick string
	// 商家id
	_sellerId string
	// 优惠券id
	_couponId string
}

// New
