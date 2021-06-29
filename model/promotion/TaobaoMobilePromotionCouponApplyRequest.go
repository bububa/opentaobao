package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取(手淘专用) API请求
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

// 初始化TaobaoMobilePromotionCouponApplyRequest对象
func NewTaobaoMobilePromotionCouponApplyRequest() *TaobaoMobilePromotionCouponApplyRequest{
    return &TaobaoMobilePromotionCouponApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionCouponApplyRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.coupon.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionCouponApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceId Setter
// 请求唯一id，问题排查
func (r *TaobaoMobilePromotionCouponApplyRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

// TraceId Getter
func (r TaobaoMobilePromotionCouponApplyRequest) GetTraceId() string {
    return r.traceId
}
// SpreadId Setter
// 传播id
func (r *TaobaoMobilePromotionCouponApplyRequest) SetSpreadId(spreadId int64) error {
    r.spreadId = spreadId
    r.Set("spread_id", spreadId)
    return nil
}

// SpreadId Getter
func (r TaobaoMobilePromotionCouponApplyRequest) GetSpreadId() int64 {
    return r.spreadId
}
// FeedId Setter
// 广播id
func (r *TaobaoMobilePromotionCouponApplyRequest) SetFeedId(feedId string) error {
    r.feedId = feedId
    r.Set("feed_id", feedId)
    return nil
}

// FeedId Getter
func (r TaobaoMobilePromotionCouponApplyRequest) GetFeedId() string {
    return r.feedId
}
// BizId Setter
// 三方活动id
func (r *TaobaoMobilePromotionCouponApplyRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

// BizId Getter
func (r TaobaoMobilePromotionCouponApplyRequest) GetBizId() string {
    return r.bizId
}
