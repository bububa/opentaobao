package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallpromotioncouponqueryAPIRequest 查询可用优惠券列表 API请求
// tmall.promotion.coupon.query
//
// 查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
type TmallpromotioncouponqueryAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerId string
	// buyer_id、buyer_nick至少填一个， 都填写以id为准
	_buyerNick string
	// 扩展字段
	_extra string
}

// NewTmallpromotioncouponqueryRequest 初始化TmallpromotioncouponqueryAPIRequest对象
func NewTmallpromotioncouponqueryRequest() *TmallpromotioncouponqueryAPIRequest {
	return &TmallpromotioncouponqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallpromotioncouponqueryAPIRequest) GetApiMethodName() string {
	return "tmall.promotion.coupon.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallpromotioncouponqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallpromotioncouponqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务类型
func (r *TmallpromotioncouponqueryAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallpromotioncouponqueryAPIRequest) GetBizType() string {
	return r._bizType
}

// SetBuyerId is BuyerId Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallpromotioncouponqueryAPIRequest) SetBuyerId(_buyerId string) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r TmallpromotioncouponqueryAPIRequest) GetBuyerId() string {
	return r._buyerId
}

// SetBuyerNick is BuyerNick Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallpromotioncouponqueryAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TmallpromotioncouponqueryAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetExtra is Extra Setter
// 扩展字段
func (r *TmallpromotioncouponqueryAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TmallpromotioncouponqueryAPIRequest) GetExtra() string {
	return r._extra
}
