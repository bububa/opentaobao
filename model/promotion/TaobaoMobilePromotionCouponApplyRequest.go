package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取(手淘专用) APIRequest
taobao.mobile.promotion.coupon.apply

优惠券领取
*/
type TaobaoMobilePromotionCouponApplyRequest struct {
    model.Params

    // 请求唯一id，问题排查
    traceId   string 

    // 传播id
    spreadId   int64 

    // 广播id
    feedId   string 

    // 三方活动id
    bizId   string 

}

func NewTaobaoMobilePromotionCouponApplyRequest() *TaobaoMobilePromotionCouponApplyRequest{
    return &TaobaoMobilePromotionCouponApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMobilePromotionCouponApplyRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.coupon.apply"
}

func (r TaobaoMobilePromotionCouponApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMobilePromotionCouponApplyRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

func (r TaobaoMobilePromotionCouponApplyRequest) GetTraceId() string {
    return r.traceId
}

func (r *TaobaoMobilePromotionCouponApplyRequest) SetSpreadId(spreadId int64) error {
    r.spreadId = spreadId
    r.Set("spread_id", spreadId)
    return nil
}

func (r TaobaoMobilePromotionCouponApplyRequest) GetSpreadId() int64 {
    return r.spreadId
}

func (r *TaobaoMobilePromotionCouponApplyRequest) SetFeedId(feedId string) error {
    r.feedId = feedId
    r.Set("feed_id", feedId)
    return nil
}

func (r TaobaoMobilePromotionCouponApplyRequest) GetFeedId() string {
    return r.feedId
}

func (r *TaobaoMobilePromotionCouponApplyRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r TaobaoMobilePromotionCouponApplyRequest) GetBizId() string {
    return r.bizId
}

