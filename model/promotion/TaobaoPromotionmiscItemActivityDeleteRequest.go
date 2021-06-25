package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除无条件单品优惠活动 APIRequest
taobao.promotionmisc.item.activity.delete

删除无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityDeleteRequest struct {
    model.Params

    // 活动id。
    activityId   int64 

}

func NewTaobaoPromotionmiscItemActivityDeleteRequest() *TaobaoPromotionmiscItemActivityDeleteRequest{
    return &TaobaoPromotionmiscItemActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscItemActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.delete"
}

func (r TaobaoPromotionmiscItemActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscItemActivityDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscItemActivityDeleteRequest) GetActivityId() int64 {
    return r.activityId
}

