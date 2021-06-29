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
type TmallPromotionCouponUseRequest struct {
    model.Params
    // 扩展字段
    extra   string
    // 业务类型
    bizType   string
    // buyer_id、buyer_nick至少填一个， 都填写以id为准
    buyerId   string
    // buyer_id、buyer_nick至少填一个， 都填写以id为准
    buyerNick   string
    // 商家id
    sellerId   string
    // 优惠券id
    couponId   string
}

// 初始化TmallPromotionCouponUseRequest对象
func NewTmallPromotionCouponUseRequest() *TmallPromotionCouponUseRequest{
    return &TmallPromotionCouponUseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotionCouponUseRequest) GetApiMethodName() string {
    return "tmall.promotion.coupon.use"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotionCouponUseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extra Setter
// 扩展字段
func (r *TmallPromotionCouponUseRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r TmallPromotionCouponUseRequest) GetExtra() string {
    return r.extra
}
// BizType Setter
// 业务类型
func (r *TmallPromotionCouponUseRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallPromotionCouponUseRequest) GetBizType() string {
    return r.bizType
}
// BuyerId Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponUseRequest) SetBuyerId(buyerId string) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

// BuyerId Getter
func (r TmallPromotionCouponUseRequest) GetBuyerId() string {
    return r.buyerId
}
// BuyerNick Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponUseRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallPromotionCouponUseRequest) GetBuyerNick() string {
    return r.buyerNick
}
// SellerId Setter
// 商家id
func (r *TmallPromotionCouponUseRequest) SetSellerId(sellerId string) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TmallPromotionCouponUseRequest) GetSellerId() string {
    return r.sellerId
}
// CouponId Setter
// 优惠券id
func (r *TmallPromotionCouponUseRequest) SetCouponId(couponId string) error {
    r.couponId = couponId
    r.Set("coupon_id", couponId)
    return nil
}

// CouponId Getter
func (r TmallPromotionCouponUseRequest) GetCouponId() string {
    return r.couponId
}
