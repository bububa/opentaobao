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
type TaobaoSingletreasureActivityUpdateAPIRequest struct {
    model.Params
    // 系统入参
    _activityInfo   *ActivityInfoCreateDto
}

// 初始化TaobaoSingletreasureActivityUpdateAPIRequest对象
func NewTaobaoSingletreasureActivityUpdateRequest() *TaobaoSingletreasureActivityUpdateAPIRequest{
    return &TaobaoSingletreasureActivityUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityUpdateAPIRequest) SetActivityInfo(_activityInfo *ActivityInfoCreateDto) error {
    r._activityInfo = _activityInfo
    r.Set("activity_info", _activityInfo)
    return nil
}

// ActivityInfo Getter
func (r TaobaoSingletreasureActivityUpdateAPIRequest) GetActivityInfo() *ActivityInfoCreateDto {
    return r._activityInfo
}
