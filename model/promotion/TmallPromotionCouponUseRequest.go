package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
券核销接口 APIRequest
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

func NewTmallPromotionCouponUseRequest() *TmallPromotionCouponUseRequest{
    return &TmallPromotionCouponUseRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPromotionCouponUseRequest) GetApiMethodName() string {
    return "tmall.promotion.coupon.use"
}

func (r TmallPromotionCouponUseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPromotionCouponUseRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r TmallPromotionCouponUseRequest) GetExtra() string {
    return r.extra
}

func (r *TmallPromotionCouponUseRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallPromotionCouponUseRequest) GetBizType() string {
    return r.bizType
}

func (r *TmallPromotionCouponUseRequest) SetBuyerId(buyerId string) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

func (r TmallPromotionCouponUseRequest) GetBuyerId() string {
    return r.buyerId
}

func (r *TmallPromotionCouponUseRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TmallPromotionCouponUseRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *TmallPromotionCouponUseRequest) SetSellerId(sellerId string) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TmallPromotionCouponUseRequest) GetSellerId() string {
    return r.sellerId
}

func (r *TmallPromotionCouponUseRequest) SetCouponId(couponId string) error {
    r.couponId = couponId
    r.Set("coupon_id", couponId)
    return nil
}

func (r TmallPromotionCouponUseRequest) GetCouponId() string {
    return r.couponId
}

