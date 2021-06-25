package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除满就送活动 APIRequest
taobao.promotionmisc.mjs.activity.delete

删除满就送活动
*/
type TaobaoPromotionmiscMjsActivityDeleteRequest struct {
    model.Params

    // 活动id。
    activityId   int64 

}

func NewTaobaoPromotionmiscMjsActivityDeleteRequest() *TaobaoPromotionmiscMjsActivityDeleteRequest{
    return &TaobaoPromotionmiscMjsActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscMjsActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.mjs.activity.delete"
}

func (r TaobaoPromotionmiscMjsActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscMjsActivityDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityDeleteRequest) GetActivityId() int64 {
    return r.activityId
}

