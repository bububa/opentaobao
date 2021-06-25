package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新关联活动有效时间 APIRequest
taobao.promotion.benefit.activity.time.update

更新关联权益的活动有效时间
*/
type TaobaoPromotionBenefitActivityTimeUpdateRequest struct {
    model.Params

    // ISV活动关联权益后获得的关联ID
    relationId   int64 

    // 活动的开始时间
    startTime   string 

    // 活动的i结束时间
    endTime   string 

}

func NewTaobaoPromotionBenefitActivityTimeUpdateRequest() *TaobaoPromotionBenefitActivityTimeUpdateRequest{
    return &TaobaoPromotionBenefitActivityTimeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionBenefitActivityTimeUpdateRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.time.update"
}

func (r TaobaoPromotionBenefitActivityTimeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionBenefitActivityTimeUpdateRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

func (r TaobaoPromotionBenefitActivityTimeUpdateRequest) GetRelationId() int64 {
    return r.relationId
}

func (r *TaobaoPromotionBenefitActivityTimeUpdateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoPromotionBenefitActivityTimeUpdateRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoPromotionBenefitActivityTimeUpdateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionBenefitActivityTimeUpdateRequest) GetEndTime() string {
    return r.endTime
}

