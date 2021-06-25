package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除通用单品优惠详情 APIRequest
taobao.promotionmisc.common.item.detail.delete

删除通用单品优惠详情。
*/
type TaobaoPromotionmiscCommonItemDetailDeleteRequest struct {
    model.Params

    // 优惠活动ID
    activityId   int64 

    // 优惠详情ID
    detailId   int64 

}

func NewTaobaoPromotionmiscCommonItemDetailDeleteRequest() *TaobaoPromotionmiscCommonItemDetailDeleteRequest{
    return &TaobaoPromotionmiscCommonItemDetailDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.detail.delete"
}

func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemDetailDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoPromotionmiscCommonItemDetailDeleteRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

func (r TaobaoPromotionmiscCommonItemDetailDeleteRequest) GetDetailId() int64 {
    return r.detailId
}

