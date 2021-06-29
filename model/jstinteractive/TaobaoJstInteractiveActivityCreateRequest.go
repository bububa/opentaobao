package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动创建接口 APIRequest
taobao.jst.interactive.activity.create

调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回
*/
type TaobaoJstInteractiveActivityCreateRequest struct {
    model.Params

    // 小程序id
    miniAppId   string 

    // 活动名称
    activityName   string 

    // 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    startTime   string 

    // 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
    endTime   string 

}

func NewTaobaoJstInteractiveActivityCreateRequest() *TaobaoJstInteractiveActivityCreateRequest{
    return &TaobaoJstInteractiveActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractiveActivityCreateRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.activity.create"
}

func (r TaobaoJstInteractiveActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstInteractiveActivityCreateRequest) SetMiniAppId(miniAppId string) error {
    r.miniAppId = miniAppId
    r.Set("mini_app_id", miniAppId)
    return nil
}

func (r TaobaoJstInteractiveActivityCreateRequest) GetMiniAppId() string {
    return r.miniAppId
}

func (r *TaobaoJstInteractiveActivityCreateRequest) SetActivityName(activityName string) error {
    r.activityName = activityName
    r.Set("activity_name", activityName)
    return nil
}

func (r TaobaoJstInteractiveActivityCreateRequest) GetActivityName() string {
    return r.activityName
}

func (r *TaobaoJstInteractiveActivityCreateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoJstInteractiveActivityCreateRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoJstInteractiveActivityCreateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoJstInteractiveActivityCreateRequest) GetEndTime() string {
    return r.endTime
}

