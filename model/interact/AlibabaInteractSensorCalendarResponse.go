package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) APIResponse
alibaba.interact.sensor.calendar

天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
*/
type AlibabaInteractSensorCalendarAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorCalendarResponse `json:"alibaba_interact_sensor_calendar_response,omitempty"`
}

type AlibabaInteractSensorCalendarResponse struct {

    // 返回结果
    Succ   bool `json:"succ,omitempty"`

}
