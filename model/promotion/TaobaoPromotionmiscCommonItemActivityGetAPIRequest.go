package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠活动 API请求
taobao.promotionmisc.common.item.activity.get

查询通用单品优惠活动。
*/
type TaobaoPromotionmiscCommonItemActivityGetAPIRequest struct {
    model.Params
    // 优惠活动ID
    _activityId   int64
}

// 初始化TaobaoPromotionmiscCommonItemActivityGetAPIRequest对象
func NewTaobaoPromotionmiscCommonItemActivityGetRequest() *TaobaoPromotionmiscCommonItemActivityGetAPIRequest{
    return &TaobaoPromotionmiscCommonItemActivityGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemActivityGetAPIRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetActivityId() int64 {
    return r._activityId
}
