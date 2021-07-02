package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotionCouponQueryAPIRequest 查询可用优惠券列表 API请求
// tmall.promotion.coupon.query
//
// 查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
type TmallPromotionCouponQueryAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerId string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerNick string
}

// NewTmallPromotionCouponQueryRequest 初始化TmallPromotionCouponQueryAPIRequest对象
func NewTmallPromotionCouponQueryRequest() *TmallPromotionCouponQueryAPIRequest {
	return &TmallPromotionCouponQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotionCouponQueryAPIRequest) GetApiMethodName() string {
	return "tmall.promotion.coupon.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotionCouponQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizType Setter
// 业务类型
func (r *TmallPromotionCouponQueryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r TmallPromotionCouponQueryAPIRequest) GetBizType() string {
	return r._bizType
}

// Set is BuyerId Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponQueryAPIRequest) SetBuyerId(_buyerId string) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// Get BuyerId Getter
func (r TmallPromotionCouponQueryAPIRequest) GetBuyerId() string {
	return r._buyerId
}

// Set is BuyerNick Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponQueryAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// Get BuyerNick Getter
func (r TmallPromotionCouponQueryAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
