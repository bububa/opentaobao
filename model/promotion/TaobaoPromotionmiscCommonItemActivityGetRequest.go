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
type TaobaoPromotionmiscCommonItemActivityGetRequest struct {
    model.Params
    // 优惠活动ID
    activityId   int64
}

// 初始化TaobaoPromotionmiscCommonItemActivityGetRequest对象
func NewTaobaoPromotionmiscCommonItemActivityGetRequest() *TaobaoPromotionmiscCommonItemActivityGetRequest{
    return &TaobaoPromotionmiscCommonItemActivityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemActivityGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemActivityGetRequest) GetActivityId() int64 {
    return r.activityId
}
