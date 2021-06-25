package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘专用单用户发放接口 APIRequest
taobao.mobile.promotion.benefit.activity.send.share

卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。
*/
type TaobaoMobilePromotionBenefitActivitySendShareRequest struct {
    model.Params

    // 权益类型    其中ALIPAY_COUPON 对应的type值是1
    benefitType   int64 

    // 权益关联的活动ID
    bizId   string 

    // 活动详情id
    detailId   int64 

    // 广播ID
    feedId   int64 

    // 关联活动id
    relationId   int64 

    // 权益发放数量
    sendCount   int64 

    // 和bizId一起使用，标记分享链路的key。
    shareKey   string 

    // 分享链路上的用户及用户是否发奖，u1_true,u2_true
    shareUsers   string 

    // 调试线索
    traceId   string 

    // 事务id
    uniqueId   string 

}

func NewTaobaoMobilePromotionBenefitActivitySendShareRequest() *TaobaoMobilePromotionBenefitActivitySendShareRequest{
    return &TaobaoMobilePromotionBenefitActivitySendShareRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetApiMethodName() string {
    return "taobao.mobile.promotion.benefit.activity.send.share"
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetBenefitType(benefitType int64) error {
    r.benefitType = benefitType
    r.Set("benefit_type", benefitType)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetBenefitType() int64 {
    return r.benefitType
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetBizId() string {
    return r.bizId
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetDetailId() int64 {
    return r.detailId
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetFeedId(feedId int64) error {
    r.feedId = feedId
    r.Set("feed_id", feedId)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetFeedId() int64 {
    return r.feedId
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetRelationId() int64 {
    return r.relationId
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetSendCount(sendCount int64) error {
    r.sendCount = sendCount
    r.Set("send_count", sendCount)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetSendCount() int64 {
    return r.sendCount
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetShareKey(shareKey string) error {
    r.shareKey = shareKey
    r.Set("share_key", shareKey)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetShareKey() string {
    return r.shareKey
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetShareUsers(shareUsers string) error {
    r.shareUsers = shareUsers
    r.Set("share_users", shareUsers)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetShareUsers() string {
    return r.shareUsers
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetTraceId() string {
    return r.traceId
}

func (r *TaobaoMobilePromotionBenefitActivitySendShareRequest) SetUniqueId(uniqueId string) error {
    r.uniqueId = uniqueId
    r.Set("unique_id", uniqueId)
    return nil
}

func (r TaobaoMobilePromotionBenefitActivitySendShareRequest) GetUniqueId() string {
    return r.uniqueId
}

