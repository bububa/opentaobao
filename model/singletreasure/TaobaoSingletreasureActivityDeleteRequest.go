package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动接口 API请求
taobao.singletreasure.activity.delete

删除优惠活动
*/
type TaobaoSingletreasureActivityDeleteRequest struct {
    model.Params
    // 活动Id
    activityId   int64
}

// 初始化TaobaoSingletreasureActivityDeleteRequest对象
func NewTaobaoSingletreasureActivityDeleteRequest() *TaobaoSingletreasureActivityDeleteRequest{
    return &TaobaoSingletreasureActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动Id
func (r *TaobaoSingletreasureActivityDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoSingletreasureActivityDeleteRequest) GetActivityId() int64 {
    return r.activityId
}
