package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动接口 API请求
taobao.singletreasure.activity.update

修改活动接口
*/
type TaobaoSingletreasureActivityUpdateRequest struct {
    model.Params
    // 系统入参
    _activityInfo   *ActivityInfoCreateDto
}

// 初始化TaobaoSingletreasureActivityUpdateRequest对象
func NewTaobaoSingletreasureActivityUpdateRequest() *TaobaoSingletreasureActivityUpdateRequest{
    return &TaobaoSingletreasureActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityUpdateRequest) SetActivityInfo(_activityInfo *ActivityInfoCreateDto) error {
    r._activityInfo = _activityInfo
    r.Set("activity_info", _activityInfo)
    return nil
}

// ActivityInfo Getter
func (r TaobaoSingletreasureActivityUpdateRequest) GetActivityInfo() *ActivityInfoCreateDto {
    return r._activityInfo
}
