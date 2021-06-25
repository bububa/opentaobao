package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动 APIRequest
taobao.promotionmisc.item.activity.get

查询无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityGetRequest struct {
    model.Params

    // 活动id。
    activityId   int64 

}

func NewTaobaoPromotionmiscItemActivityGetRequest() *TaobaoPromotionmiscItemActivityGetRequest{
    return &TaobaoPromotionmiscItemActivityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscItemActivityGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.get"
}

func (r TaobaoPromotionmiscItemActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscItemActivityGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscItemActivityGetRequest) GetActivityId() int64 {
    return r.activityId
}

