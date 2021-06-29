package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除满就送活动 API请求
taobao.promotionmisc.mjs.activity.delete

删除满就送活动
*/
type TaobaoPromotionmiscMjsActivityDeleteRequest struct {
    model.Params
    // 活动id。
    _activityId   int64
}

// 初始化TaobaoPromotionmiscMjsActivityDeleteRequest对象
func NewTaobaoPromotionmiscMjsActivityDeleteRequest() *TaobaoPromotionmiscMjsActivityDeleteRequest{
    return &TaobaoPromotionmiscMjsActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.mjs.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscMjsActivityDeleteRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscMjsActivityDeleteRequest) GetActivityId() int64 {
    return r._activityId
}
