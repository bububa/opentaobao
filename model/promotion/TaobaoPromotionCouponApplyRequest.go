package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取 API请求
taobao.promotion.coupon.apply

优惠券领取
*/
type TaobaoPromotionCouponApplyRequest struct {
    model.Params
    // 卖家id
    _sellerId   string
    // 传播id
    _spreadId   string
}

// 初始化TaobaoPromotionCouponApplyRequest对象
func NewTaobaoPromotionCouponApplyRequest() *TaobaoPromotionCouponApplyRequest{
    return &TaobaoPromotionCouponApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponApplyRequest) GetApiMethodName() string {
    return "taobao.promotion.coupon.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 卖家id
func (r *TaobaoPromotionCouponApplyRequest) SetSellerId(_sellerId string) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoPromotionCouponApplyRequest) GetSellerId() string {
    return r._sellerId
}
// SpreadId Setter
// 传播id
func (r *TaobaoPromotionCouponApplyRequest) SetSpreadId(_spreadId string) error {
    r._spreadId = _spreadId
    r.Set("spread_id", _spreadId)
    return nil
}

// SpreadId Getter
func (r TaobaoPromotionCouponApplyRequest) GetSpreadId() string {
    return r._spreadId
}
