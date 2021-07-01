package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动修改接口 API请求
taobao.jst.interactive.activity.update

互动任务活动修改接口
*/
type TaobaoJstInteractiveActivityUpdateAPIRequest struct {
    model.Params
    // 小程序id
    _miniAppId   string
    // 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    _startTime   string
    // 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    _endTime   string
    // 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
    _status   int64
}

// 初始化TaobaoJstInteractiveActivityUpdateAPIRequest对象
func NewTaobaoJstInteractiveActivityUpdateRequest() *TaobaoJstInteractiveActivityUpdateAPIRequest{
    return &TaobaoJstInteractiveActivityUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetMiniAppId(_miniAppId string) error {
    r._miniAppId = _miniAppId
    r.Set("mini_app_id", _miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetMiniAppId() string {
    return r._miniAppId
}
// StartTime Setter
// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetEndTime() string {
    return r._endTime
}
// Status Setter
// 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetStatus() int64 {
    return r._status
}
