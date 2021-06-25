package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取 APIRequest
taobao.promotion.coupon.apply

优惠券领取
*/
type TaobaoPromotionCouponApplyRequest struct {
    model.Params

    // 卖家id
    sellerId   string 

    // 传播id
    spreadId   string 

}

func NewTaobaoPromotionCouponApplyRequest() *TaobaoPromotionCouponApplyRequest{
    return &TaobaoPromotionCouponApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionCouponApplyRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.apply"
}

func (r TaobaoPromotionCouponApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionCouponApplyRequest) SetSellerId(sellerId string) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TaobaoPromotionCouponApplyRequest) GetSellerId() string {
    return r.sellerId
}

func (r *TaobaoPromotionCouponApplyRequest) SetSpreadId(spreadId string) error {
    r.spreadId = spreadId
    r.Set("spread_id", spreadId)
    return nil
}

func (r TaobaoPromotionCouponApplyRequest) GetSpreadId() string {
    return r.spreadId
}

