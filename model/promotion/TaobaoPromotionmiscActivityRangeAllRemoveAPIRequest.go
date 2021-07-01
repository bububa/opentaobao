package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
清空活动参与的商品 API请求
taobao.promotionmisc.activity.range.all.remove

清空活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest struct {
    model.Params
    // 活动id。
    _activityId   int64
}

// 初始化TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest对象
func NewTaobaoPromotionmiscActivityRangeAllRemoveRequest() *TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest{
    return &TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.all.remove"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) GetActivityId() int64 {
    return r._activityId
}
