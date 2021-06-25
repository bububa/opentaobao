package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
去除活动参与的商品 APIRequest
taobao.promotionmisc.activity.range.remove

去除活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeRemoveRequest struct {
    model.Params

    // 活动id。
    activityId   int64 

    // 商品id,多个id用逗号隔开。
    ids   string 

}

func NewTaobaoPromotionmiscActivityRangeRemoveRequest() *TaobaoPromotionmiscActivityRangeRemoveRequest{
    return &TaobaoPromotionmiscActivityRangeRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.remove"
}

func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscActivityRangeRemoveRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoPromotionmiscActivityRangeRemoveRequest) SetIds(ids string) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TaobaoPromotionmiscActivityRangeRemoveRequest) GetIds() string {
    return r.ids
}

