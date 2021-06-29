package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动接口 APIRequest
taobao.singletreasure.activity.update

修改活动接口
*/
type TaobaoSingletreasureActivityUpdateRequest struct {
    model.Params

    // 系统入参
    activityInfo   *ActivityInfoCreateDto 

}

func NewTaobaoSingletreasureActivityUpdateRequest() *TaobaoSingletreasureActivityUpdateRequest{
    return &TaobaoSingletreasureActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.update"
}

func (r TaobaoSingletreasureActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityUpdateRequest) SetActivityInfo(activityInfo *ActivityInfoCreateDto) error {
    r.activityInfo = activityInfo
    r.Set("activity_info", activityInfo)
    return nil
}

func (r TaobaoSingletreasureActivityUpdateRequest) GetActivityInfo() *ActivityInfoCreateDto {
    return r.activityInfo
}

