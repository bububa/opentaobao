package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠活动 APIRequest
taobao.promotionmisc.common.item.activity.get

查询通用单品优惠活动。
*/
type TaobaoPromotionmiscCommonItemActivityGetRequest struct {
    model.Params

    // 优惠活动ID
    activityId   int64 

}

func NewTaobaoPromotionmiscCommonItemActivityGetRequest() *TaobaoPromotionmiscCommonItemActivityGetRequest{
    return &TaobaoPromotionmiscCommonItemActivityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemActivityGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.get"
}

func (r TaobaoPromotionmiscCommonItemActivityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemActivityGetRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityGetRequest) GetActivityId() int64 {
    return r.activityId
}

