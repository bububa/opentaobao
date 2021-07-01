package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新关联活动有效时间 API请求
taobao.promotion.benefit.activity.time.update

更新关联权益的活动有效时间
*/
type TaobaoPromotionBenefitActivityTimeUpdateAPIRequest struct {
    model.Params
    // ISV活动关联权益后获得的关联ID
    _relationId   int64
    // 活动的开始时间
    _startTime   string
    // 活动的i结束时间
    _endTime   string
}

// 初始化TaobaoPromotionBenefitActivityTimeUpdateAPIRequest对象
func NewTaobaoPromotionBenefitActivityTimeUpdateRequest() *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest{
    return &TaobaoPromotionBenefitActivityTimeUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.time.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RelationId Setter
// ISV活动关联权益后获得的关联ID
func (r *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) SetRelationId(_relationId int64) error {
    r._relationId = _relationId
    r.Set("relation_id", _relationId)
    return nil
}

// RelationId Getter
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetRelationId() int64 {
    return r._relationId
}
// StartTime Setter
// 活动的开始时间
func (r *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 活动的i结束时间
func (r *TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionBenefitActivityTimeUpdateAPIRequest) GetEndTime() string {
    return r._endTime
}
