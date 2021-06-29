package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动创建接口 API请求
taobao.jst.interactive.activity.create

调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回
*/
type TaobaoJstInteractiveActivityCreateRequest struct {
    model.Params
    // 小程序id
    _miniAppId   string
    // 活动名称
    _activityName   string
    // 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    _startTime   string
    // 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    _endTime   string
}

// 初始化TaobaoJstInteractiveActivityCreateRequest对象
func NewTaobaoJstInteractiveActivityCreateRequest() *TaobaoJstInteractiveActivityCreateRequest{
    return &TaobaoJstInteractiveActivityCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveActivityCreateRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.activity.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveActivityCreateRequest) SetMiniAppId(_miniAppId string) error {
    r._miniAppId = _miniAppId
    r.Set("mini_app_id", _miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveActivityCreateRequest) GetMiniAppId() string {
    return r._miniAppId
}
// ActivityName Setter
// 活动名称
func (r *TaobaoJstInteractiveActivityCreateRequest) SetActivityName(_activityName string) error {
    r._activityName = _activityName
    r.Set("activity_name", _activityName)
    return nil
}

// ActivityName Getter
func (r TaobaoJstInteractiveActivityCreateRequest) GetActivityName() string {
    return r._activityName
}
// StartTime Setter
// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityCreateRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoJstInteractiveActivityCreateRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityCreateRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoJstInteractiveActivityCreateRequest) GetEndTime() string {
    return r._endTime
}
