package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动创建接口 API请求
taobao.singletreasure.activity.create

创建优惠活动
*/
type TaobaoSingletreasureActivityCreateRequest struct {
    model.Params
    // 系统入参
    activityInfo   *ActivityInfoCreateDto
}

// 初始化TaobaoSingletreasureActivityCreateRequest对象
func NewTaobaoSingletreasureActivityCreateRequest() *TaobaoSingletreasureActivityCreateRequest{
    return &TaobaoSingletreasureActivityCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityCreateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityCreateRequest) SetActivityInfo(activityInfo *ActivityInfoCreateDto) error {
    r.activityInfo = activityInfo
    r.Set("activity_info", activityInfo)
    return nil
}

// ActivityInfo Getter
func (r TaobaoSingletreasureActivityCreateRequest) GetActivityInfo() *ActivityInfoCreateDto {
    return r.activityInfo
}
