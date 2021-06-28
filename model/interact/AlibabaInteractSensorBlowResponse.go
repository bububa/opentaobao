package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
吹气 APIResponse
alibaba.interact.sensor.blow

客户端吹气
*/
type AlibabaInteractSensorBlowAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorBlowResponse `json:"alibaba_interact_sensor_blow_response,omitempty"` 
    AlibabaInteractSensorBlowResponse
}

/* model for simplify = false
type AlibabaInteractSensorBlowResponse struct {

    // return=0 表示成功
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorBlowResponse struct {

    // return=0 表示成功
    Result   string `json:"result,omitempty"`

}
