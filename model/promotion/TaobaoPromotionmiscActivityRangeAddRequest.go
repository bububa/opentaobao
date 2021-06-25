package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加活动参与的商品 APIRequest
taobao.promotionmisc.activity.range.add

增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
*/
type TaobaoPromotionmiscActivityRangeAddRequest struct {
    model.Params

    // 活动id。
    activityId   int64 

    // 商品id,多个id用逗号隔开，一次不超过50个。
    ids   string 

}

func NewTaobaoPromotionmiscActivityRangeAddRequest() *TaobaoPromotionmiscActivityRangeAddRequest{
    return &TaobaoPromotionmiscActivityRangeAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscActivityRangeAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.activity.range.add"
}

func (r TaobaoPromotionmiscActivityRangeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscActivityRangeAddRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscActivityRangeAddRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoPromotionmiscActivityRangeAddRequest) SetIds(ids string) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TaobaoPromotionmiscActivityRangeAddRequest) GetIds() string {
    return r.ids
}

