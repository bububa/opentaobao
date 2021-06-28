package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
摇一摇 APIResponse
alibaba.interact.sensor.shake

摇一摇
*/
type AlibabaInteractSensorShakeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_shake_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0
    
    Result   string `json:"result,omitempty" xml:"