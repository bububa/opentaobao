package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动创建接口 APIRequest
taobao.singletreasure.activity.create

创建优惠活动
*/
type TaobaoSingletreasureActivityCreateRequest struct {
    model.Params

    // 系统入参
    activityInfo   *ActivityInfoCreateDto 

}

func NewTaobaoSingletreasureActivityCreateRequest() *TaobaoSingletreasureActivityCreateRequest{
    return &TaobaoSingletreasureActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityCreateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.create"
}

func (r TaobaoSingletreasureActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityCreateRequest) SetActivityInfo(activityInfo *ActivityInfoCreateDto) error {
    r.activityInfo = activityInfo
    r.Set("activity_info", activityInfo)
    return nil
}

func (r TaobaoSingletreasureActivityCreateRequest) GetActivityInfo() *ActivityInfoCreateDto {
    return r.activityInfo
}

