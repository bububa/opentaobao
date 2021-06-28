package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
重力感应 APIResponse
alibaba.interact.sensor.gravity

native获取重力感应
*/
type AlibabaInteractSensorGravityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorGravityResponse `json:"alibaba_interact_sensor_gravity_response,omitempty"` 
    AlibabaInteractSensorGravityResponse
}

/* model for simplify = false
type AlibabaInteractSensorGravityResponse struct {

    // 0表示成功呢
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorGravityResponse struct {

    // 0表示成功呢
    Result   string `json:"result,omitempty"`

}
