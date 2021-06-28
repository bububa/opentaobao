package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
重力感应 APIResponse
alibaba.interact.sensor.gravity

native获取重力感应
*/
type AlibabaInteractSensorGravityAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_gravity_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0表示成功呢
    
    Result   string `json:"result,omitempty" xml:"