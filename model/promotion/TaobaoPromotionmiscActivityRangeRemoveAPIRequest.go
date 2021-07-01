package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
去除活动参与的商品 API请求
taobao.promotionmisc.activity.range.remove

去除活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeRemoveAPIRequest struct {
    model.Params
    // 活动id。
    _activityId   int64
    // 商品id,多个id用逗号隔开。
    _ids   string
}

// 初始化TaobaoPromotionmiscActivityRangeRemoveAPIRequest对象
func NewTaobaoPromotionmiscActivityRangeRemoveRequest() *TaobaoPromotionmiscActivityRangeRemoveAPIRequest{
    return &TaobaoPromotionmiscActivityRangeRemoveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.remove"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeRemoveAPIRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetActivityId() int64 {
    return r._activityId
}
// Ids Setter
// 商品id,多个id用逗号隔开。
func (r *TaobaoPromotionmiscActivityRangeRemoveAPIRequest) SetIds(_ids string) error {
    r._ids = _ids
    r.Set("ids", _ids)
    return nil
}

// Ids Getter
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetIds() string {
    return r._ids
}
