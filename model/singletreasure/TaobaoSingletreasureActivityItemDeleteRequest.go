package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品优惠接口 API请求
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

// 初始化TaobaoSingletreasureActivityItemDeleteRequest对象
func NewTaobaoSingletreasureActivityItemDeleteRequest() *TaobaoSingletreasureActivityItemDeleteRequest{
    return &TaobaoSingletreasureActivityItemDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemDeleteRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id
func (r *TaobaoSingletreasureActivityItemDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoSingletreasureActivityItemDeleteRequest) GetActivityId() int64 {
    return r.activityId
}
// ItemId Setter
// 商品Id
func (r *TaobaoSingletreasureActivityItemDeleteRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSingletreasureActivityItemDeleteRequest) GetItemId() int64 {
    return r.itemId
}
