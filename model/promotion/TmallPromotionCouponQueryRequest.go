package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询可用优惠券列表 API请求
tmall.promotion.coupon.query

查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
*/
type TmallPromotionCouponQueryRequest struct {
    model.Params
    // 业务类型
    bizType   string
    // buyer_id、buyer_nick至少填一个， 都填写以id为准
    buyerId   string
    // buyer_id、buyer_nick至少填一个， 都填写以id为准
    buyerNick   string
}

// 初始化TmallPromotionCouponQueryRequest对象
func NewTmallPromotionCouponQueryRequest() *TmallPromotionCouponQueryRequest{
    return &TmallPromotionCouponQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotionCouponQueryRequest) GetApiMethodName() string {
    return "tmall.promotion.coupon.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotionCouponQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TmallPromotionCouponQueryRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallPromotionCouponQueryRequest) GetBizType() string {
    return r.bizType
}
// BuyerId Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponQueryRequest) SetBuyerId(buyerId string) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

// BuyerId Getter
func (r TmallPromotionCouponQueryRequest) GetBuyerId() string {
    return r.buyerId
}
// BuyerNick Setter
// buyer_id、buyer_nick至少填一个， 都填写以id为准
func (r *TmallPromotionCouponQueryRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallPromotionCouponQueryRequest) GetBuyerNick() string {
    return r.buyerNick
}
