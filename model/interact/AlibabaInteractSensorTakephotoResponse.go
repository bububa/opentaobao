package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
takePhoto APIResponse
alibaba.interact.sensor.takephoto

客户端takePhoto
*/
type AlibabaInteractSensorTakephotoAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorTakephotoResponse `json:"alibaba_interact_sensor_takephoto_response,omitempty"` 
    AlibabaInteractSensorTakephotoResponse
}

/* model for simplify = false
type AlibabaInteractSensorTakephotoResponse struct {

    // return=0表示成功
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorTakephotoResponse struct {

    // return=0表示成功
    Result   string `json:"result,omitempty"`

}
