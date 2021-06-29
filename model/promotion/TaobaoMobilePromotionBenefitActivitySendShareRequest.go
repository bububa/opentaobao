package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘专用单用户发放接口 API请求
taobao.mobile.promotion.benefit.activity.send.share

卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。
*/
type TaobaoMobilePromotionBenefitActivitySendShareRequest struct {
    model.Params
    // 权益类型    其中ALIPAY_COUPON 对应的type值是1
    _benefitType   int64
    // 权益关联的活动ID
    _bizId   string
    // 活动详情id
    _detailId   int64
    // 广播ID
    _feedId   int64
    // 关联活动id
    _relationId   int64
    // 权益发放数量
    _sendCount   int64
    // 和bizId一起使用，标记分享链路的key。
    _shareKey   string
    // 分享链路上的用户及用户是否发奖，u1_true,u2_true
    _shareUsers   string
    // 调试线索
    _traceId   string
    // 事务id
    _uniqueId   string
}

// 初始化TaobaoMobilePromotionBenefitActivitySendShareRequest对象
func NewTaobaoMobilePromotionBenefitActivitySendShareRequest() *TaobaoMobilePromotionBenefitActivitySendShareRequest{
    return &TaobaoMobilePromotionBenefitActivitySendShareRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.benefit.activity.send.share"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BenefitType Setter
// 权益类型    其中ALIPAY_COUPON 对应的type值是1
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetBenefitType(_benefitType int64) error {
    r._benefitType = _benefitType
    r.Set("benefit_type", _benefitType)
    return nil
}

// BenefitType Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetBenefitType() int64 {
    return r._benefitType
}
// BizId Setter
// 权益关联的活动ID
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetBizId(_bizId string) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetBizId() string {
    return r._bizId
}
// DetailId Setter
// 活动详情id
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetDetailId(_detailId int64) error {
    r._detailId = _detailId
    r.Set("detail_id", _detailId)
    return nil
}

// DetailId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetDetailId() int64 {
    return r._detailId
}
// FeedId Setter
// 广播ID
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetFeedId(_feedId int64) error {
    r._feedId = _feedId
    r.Set("feed_id", _feedId)
    return nil
}

// FeedId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetFeedId() int64 {
    return r._feedId
}
// RelationId Setter
// 关联活动id
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetRelationId(_relationId int64) error {
    r._relationId = _relationId
    r.Set("relation_id", _relationId)
    return nil
}

// RelationId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetRelationId() int64 {
    return r._relationId
}
// SendCount Setter
// 权益发放数量
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetSendCount(_sendCount int64) error {
    r._sendCount = _sendCount
    r.Set("send_count", _sendCount)
    return nil
}

// SendCount Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetSendCount() int64 {
    return r._sendCount
}
// ShareKey Setter
// 和bizId一起使用，标记分享链路的key。
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetShareKey(_shareKey string) error {
    r._shareKey = _shareKey
    r.Set("share_key", _shareKey)
    return nil
}

// ShareKey Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetShareKey() string {
    return r._shareKey
}
// ShareUsers Setter
// 分享链路上的用户及用户是否发奖，u1_true,u2_true
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetShareUsers(_shareUsers string) error {
    r._shareUsers = _shareUsers
    r.Set("share_users", _shareUsers)
    return nil
}

// ShareUsers Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetShareUsers() string {
    return r._shareUsers
}
// TraceId Setter
// 调试线索
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetTraceId() string {
    return r._traceId
}
// UniqueId Setter
// 事务id
func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetUniqueId(_uniqueId string) error {
    r._uniqueId = _uniqueId
    r.Set("unique_id", _uniqueId)
    return nil
}

// UniqueId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetUniqueId() string {
    return r._uniqueId
}
