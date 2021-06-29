package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.delete

删除无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityDeleteRequest struct {
    model.Params
    // 活动id。
    activityId   int64
}

// 初始化TaobaoPromotionmiscItemActivityDeleteRequest对象
func NewTaobaoPromotionmiscItemActivityDeleteRequest() *TaobaoPromotionmiscItemActivityDeleteRequest{
    return &TaobaoPromotionmiscItemActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscItemActivityDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscItemActivityDeleteRequest) GetActivityId() int64 {
    return r.activityId
}
