package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动 API请求
taobao.promotionmisc.mjs.activity.get

查询满就送活动
*/
type TaobaoPromotionmiscMjsActivityGetRequest struct {
    model.Params
    // 活动id。
    _activityId   int64
}

// 初始化TaobaoPromotionmiscMjsActivityGetRequest对象
func NewTaobaoPromotionmiscMjsActivityGetRequest() *TaobaoPromotionmiscMjsActivityGetRequest{
    return &TaobaoPromotionmiscMjsActivityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.mjs.activity.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscMjsActivityGetRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscMjsActivityGetRequest) GetActivityId() int64 {
    return r._activityId
}
