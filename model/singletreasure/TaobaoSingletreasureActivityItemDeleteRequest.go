package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品优惠接口 APIRequest
taobao.singletreasure.activity.item.delete

删除单品优惠接口
*/
type TaobaoSingletreasureActivityItemDeleteRequest struct {
    model.Params

    // 活动id
    activityId   int64 

    // 商品Id
    itemId   int64 

}

func NewTaobaoSingletreasureActivityItemDeleteRequest() *TaobaoSingletreasureActivityItemDeleteRequest{
    return &TaobaoSingletreasureActivityItemDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityItemDeleteRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.delete"
}

func (r TaobaoSingletreasureActivityItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityItemDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoSingletreasureActivityItemDeleteRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *TaobaoSingletreasureActivityItemDeleteRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoSingletreasureActivityItemDeleteRequest) GetItemId() int64 {
    return r.itemId
}

