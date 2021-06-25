package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动参与的商品 APIRequest
taobao.promotionmisc.activity.range.list.get

查询活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeListGetRequest struct {
    model.Params

    // 活动id
    activityId   int64 

}

func NewTaobaoPromotionmiscActivityRangeListGetRequest() *TaobaoPromotionmiscActivityRangeListGetRequest{
    return &TaobaoPromotionmiscActivityRangeListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscActivityRangeListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.list.get"
}

func (r TaobaoPromotionmiscActivityRangeListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscActivityRangeListGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscActivityRangeListGetRequest) GetActivityId() int64 {
    return r.activityId
}

