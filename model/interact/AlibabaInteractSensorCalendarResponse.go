package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) APIResponse
alibaba.interact.sensor.calendar

天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
*/
type AlibabaInteractSensorCalendarAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_calendar_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Succ   bool `json:"succ,omitempty" xml:"