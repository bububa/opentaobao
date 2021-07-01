package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMobilePromotionCouponApplyAPIRequest
优惠券领取(手淘专用) API请求
taobao.mobile.promotion.coupon.apply

优惠券领取 */
type TaobaoMobilePromotionCouponApplyAPIRequest struct {
	model.Params
	// 请求唯一id，问题排查
	_traceId string
	// 传播id
	_spreadId int64
	// 广播id
	_feedId string
	// 三方活动id
	_bizId string
}

// New
