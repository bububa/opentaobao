package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
gmedia APIResponse
alibaba.interact.sensor.gmedia

媒体功能
*/
type AlibabaInteractSensorGmediaAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorGmediaResponse `json:"alibaba_interact_sensor_gmedia_response,omitempty"` 
    AlibabaInteractSensorGmediaResponse
}

/* model for simplify = false
type AlibabaInteractSensorGmediaResponse struct {

    // result=0 表示成功
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorGmediaResponse struct {

    // result=0 表示成功
    Result   string `json:"result,omitempty"`

}
