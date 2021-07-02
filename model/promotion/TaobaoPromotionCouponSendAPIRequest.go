package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponSendAPIRequest 店铺优惠券发放接口 API请求
// taobao.promotion.coupon.send
//
// 通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的会员），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了
type TaobaoPromotionCouponSendAPIRequest struct {
	model.Params
	// 优惠券的id
	_couponId int64
	// 买家昵称用半角','号分割
	_buyerNick []string
}

// NewTaobaoPromotionCouponSendRequest 初始化TaobaoPromotionCouponSendAPIRequest对象
func NewTaobaoPromotionCouponSendRequest() *TaobaoPromotionCouponSendAPIRequest {
	return &TaobaoPromotionCouponSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponSendAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCouponId is CouponId Setter
// 优惠券的id
func (r *TaobaoPromotionCouponSendAPIRequest) SetCouponId(_couponId int64) error {
	r._couponId = _couponId
	r.Set("coupon_id", _couponId)
	return nil
}

// GetCouponId CouponId Getter
func (r TaobaoPromotionCouponSendAPIRequest) GetCouponId() int64 {
	return r._couponId
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称用半角','号分割
func (r *TaobaoPromotionCouponSendAPIRequest) SetBuyerNick(_buyerNick []string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoPromotionCouponSendAPIRequest) GetBuyerNick() []string {
	return r._buyerNick
}
