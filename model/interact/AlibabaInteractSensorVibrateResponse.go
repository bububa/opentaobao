package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端震动 APIResponse
alibaba.interact.sensor.vibrate

客户端震动
*/
type AlibabaInteractSensorVibrateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_vibrate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0表示成功
    
    Result   string `json:"result,omitempty" xml:"