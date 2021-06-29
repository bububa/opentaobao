package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除通用单品优惠活动 API请求
taobao.promotionmisc.common.item.activity.delete

删除通用单品优惠活动。
*/
type TaobaoPromotionmiscCommonItemActivityDeleteRequest struct {
    model.Params
    // 优惠活动ID
    activityId   int64
}

// 初始化TaobaoPromotionmiscCommonItemActivityDeleteRequest对象
func NewTaobaoPromotionmiscCommonItemActivityDeleteRequest() *TaobaoPromotionmiscCommonItemActivityDeleteRequest{
    return &TaobaoPromotionmiscCommonItemActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemActivityDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscCommonItemActivityDeleteRequest) GetActivityId() int64 {
    return r.activityId
}
