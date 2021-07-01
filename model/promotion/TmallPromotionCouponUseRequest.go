package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
券核销接口 API请求
tmall.promotion.coupon.use

核销用户的一张优惠券，返回核销结果
*/
type TmallPromotionCouponUseAPIRequest struct {
    model.Params
    // 扩展字段
    _extra   string
    // 业务类型
    _bizType   string
    // buyer_id、buyer_nick至少填一个， 都填写以id为准
    _buyerId   string
    // buyer_id、buyer_nick至少填一个， 都填写以id为准
    _buyerNick   string
    // 商家id
    _sellerId   string
    // 优惠券id
    _couponId   string
}

// 初始化TmallPromotionCouponUseAPIRequest对象
func NewTmallPromotionCouponUseRequest() *TmallPromotionCouponUseAPIRequest{
    return &TmallPromotionCouponUseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotionCouponUseAPIRequest) GetApiMethodName() string {
    return "tmall.promotion.coupon.use"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotionCouponUseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extra Setter
// 扩展字段
func (r *TmallPromotionCouponUseAPIRequest) SetExtra(_extra string) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r TmallPromotionCouponUseAPIRequest) GetExtra() string {
    return r._extra
}
// BizType Setter
// 业务类型
func (r *TmallPromotionCouponUseAPIRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallPromotionCouponUseAPIRequest) GetBizType() string {
    return r._bizType
}
// BuyerId Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponUseAPIRequest) SetBuyerId(_buyerId string) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r TmallPromotionCouponUseAPIRequest) GetBuyerId() string {
    return r._buyerId
}
// BuyerNick Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponUseAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallPromotionCouponUseAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
// SellerId Setter
// 商家id
func (r *TmallPromotionCouponUseAPIRequest) SetSellerId(_sellerId string) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TmallPromotionCouponUseAPIRequest) GetSellerId() string {
    return r._sellerId
}
// CouponId Setter
// 优惠券id
func (r *TmallPromotionCouponUseAPIRequest) SetCouponId(_couponId string) error {
    r._couponId = _couponId
    r.Set("coupon_id", _couponId)
    return nil
}

// CouponId Getter
func (r TmallPromotionCouponUseAPIRequest) GetCouponId() string {
    return r._couponId
}
