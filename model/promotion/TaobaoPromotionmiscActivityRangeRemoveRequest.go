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
type TaobaoPromotionmiscActivityRangeRemoveRequest struct {
    model.Params
    // 活动id。
    activityId   int64
    // 商品id,多个id用逗号隔开。
    ids   string
}

// 初始化TaobaoPromotionmiscActivityRangeRemoveRequest对象
func NewTaobaoPromotionmiscActivityRangeRemoveRequest() *TaobaoPromotionmiscActivityRangeRemoveRequest{
    return &TaobaoPromotionmiscActivityRangeRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.remove"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeRemoveRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetActivityId() int64 {
    return r.activityId
}
// Ids Setter
// 商品id,多个id用逗号隔开。
func (r *TaobaoPromotionmiscActivityRangeRemoveRequest) SetIds(ids string) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

// Ids Getter
func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetIds() string {
    return r.ids
}
