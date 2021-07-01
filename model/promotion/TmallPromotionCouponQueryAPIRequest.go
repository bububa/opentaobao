package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotionCouponQueryAPIRequest
查询可用优惠券列表 API请求
tmall.promotion.coupon.query

查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick */
type TmallPromotionCouponQueryAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerId string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerNick string
}

// New
