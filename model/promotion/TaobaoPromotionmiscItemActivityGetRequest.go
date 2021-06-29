package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.get

查询无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityGetRequest struct {
    model.Params
    // 活动id。
    _activityId   int64
}

// 初始化TaobaoPromotionmiscItemActivityGetRequest对象
func NewTaobaoPromotionmiscItemActivityGetRequest() *TaobaoPromotionmiscItemActivityGetRequest{
    return &TaobaoPromotionmiscItemActivityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscItemActivityGetRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscItemActivityGetRequest) GetActivityId() int64 {
    return r._activityId
}
