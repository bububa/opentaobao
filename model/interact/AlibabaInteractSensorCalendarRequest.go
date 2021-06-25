package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) APIRequest
alibaba.interact.sensor.calendar

天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
*/
type AlibabaInteractSensorCalendarRequest struct {
    model.Params

}

func NewAlibabaInteractSensorCalendarRequest() *AlibabaInteractSensorCalendarRequest{
    return &AlibabaInteractSensorCalendarRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorCalendarRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.calendar"
}

func (r AlibabaInteractSensorCalendarRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


