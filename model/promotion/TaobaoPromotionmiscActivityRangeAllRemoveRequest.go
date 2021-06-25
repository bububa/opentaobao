package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
清空活动参与的商品 APIRequest
taobao.promotionmisc.activity.range.all.remove

清空活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeAllRemoveRequest struct {
    model.Params

    // 活动id。
    activityId   int64 

}

func NewTaobaoPromotionmiscActivityRangeAllRemoveRequest() *TaobaoPromotionmiscActivityRangeAllRemoveRequest{
    return &TaobaoPromotionmiscActivityRangeAllRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscActivityRangeAllRemoveRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.all.remove"
}

func (r TaobaoPromotionmiscActivityRangeAllRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscActivityRangeAllRemoveRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscActivityRangeAllRemoveRequest) GetActivityId() int64 {
    return r.activityId
}

