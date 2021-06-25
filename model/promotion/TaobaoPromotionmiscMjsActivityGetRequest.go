package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动 APIRequest
taobao.promotionmisc.mjs.activity.get

查询满就送活动
*/
type TaobaoPromotionmiscMjsActivityGetRequest struct {
    model.Params

    // 活动id。
    activityId   int64 

}

func NewTaobaoPromotionmiscMjsActivityGetRequest() *TaobaoPromotionmiscMjsActivityGetRequest{
    return &TaobaoPromotionmiscMjsActivityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscMjsActivityGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.mjs.activity.get"
}

func (r TaobaoPromotionmiscMjsActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscMjsActivityGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityGetRequest) GetActivityId() int64 {
    return r.activityId
}

