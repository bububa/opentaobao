package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动接口 APIRequest
taobao.singletreasure.activity.delete

删除优惠活动
*/
type TaobaoSingletreasureActivityDeleteRequest struct {
    model.Params

    // 活动Id
    activityId   int64 

}

func NewTaobaoSingletreasureActivityDeleteRequest() *TaobaoSingletreasureActivityDeleteRequest{
    return &TaobaoSingletreasureActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityDeleteRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.delete"
}

func (r TaobaoSingletreasureActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityDeleteRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoSingletreasureActivityDeleteRequest) GetActivityId() int64 {
    return r.activityId
}

