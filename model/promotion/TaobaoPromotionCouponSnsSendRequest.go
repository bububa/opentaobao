package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘粉丝店铺优惠券发放接口 API请求
taobao.promotion.coupon.sns.send

通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的微淘粉丝），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了
*/
type TaobaoPromotionCouponSnsSendRequest struct {
    model.Params
    // 优惠券的id
    couponId   int64
    // 买家昵称用半角','号分割
    buyerNick   []string
}

// 初始化TaobaoPromotionCouponSnsSendRequest对象
func NewTaobaoPromotionCouponSnsSendRequest() *TaobaoPromotionCouponSnsSendRequest{
    return &TaobaoPromotionCouponSnsSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponSnsSendRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.sns.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponSnsSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CouponId Setter
// 优惠券的id
func (r *TaobaoPromotionCouponSnsSendRequest) SetCouponId(couponId int64) error {
    r.couponId = couponId
    r.Set("coupon_id", couponId)
    return nil
}

// CouponId Getter
func (r TaobaoPromotionCouponSnsSendRequest) GetCouponId() int64 {
    return r.couponId
}
// BuyerNick Setter
// 买家昵称用半角','号分割
func (r *TaobaoPromotionCouponSnsSendRequest) SetBuyerNick(buyerNick []string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoPromotionCouponSnsSendRequest) GetBuyerNick() []string {
    return r.buyerNick
}
