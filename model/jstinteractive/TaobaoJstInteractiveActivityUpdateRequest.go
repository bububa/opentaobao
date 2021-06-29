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
type TaobaoJstInteractiveActivityUpdateRequest struct {
    model.Params
    // 小程序id
    miniAppId   string
    // 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    startTime   string
    // 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    endTime   string
    // 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
    status   int64
}

// 初始化TaobaoJstInteractiveActivityUpdateRequest对象
func NewTaobaoJstInteractiveActivityUpdateRequest() *TaobaoJstInteractiveActivityUpdateRequest{
    return &TaobaoJstInteractiveActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveActivityUpdateRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

// MiniAppId Getter
func (r TaobaoJstInteractiveActivityUpdateRequest) GetMiniAppId() string {
    return r.miniAppId
}
// StartTime Setter
// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityUpdateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoJstInteractiveActivityUpdateRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityUpdateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoJstInteractiveActivityUpdateRequest) GetEndTime() string {
    return r.endTime
}
// Status Setter
// 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
func (r *TaobaoJstInteractiveActivityUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoJstInteractiveActivityUpdateRequest) GetStatus() int64 {
    return r.status
}
