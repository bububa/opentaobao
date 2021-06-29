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
type TaobaoPromotionmiscActivityRangeAllRemoveRequest struct {
    model.Params
    // 活动id。
    activityId   int64
}

// 初始化TaobaoPromotionmiscActivityRangeAllRemoveRequest对象
func NewTaobaoPromotionmiscActivityRangeAllRemoveRequest() *TaobaoPromotionmiscActivityRangeAllRemoveRequest{
    return &TaobaoPromotionmiscActivityRangeAllRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeAllRemoveRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.all.remove"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeAllRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeAllRemoveRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeAllRemoveRequest) GetActivityId() int64 {
    return r.activityId
}
