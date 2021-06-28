package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
canvas工具包 APIResponse
alibaba.interact.sensor.gutil

canvas工具包
*/
type AlibabaInteractSensorGutilAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorGutilResponse `json:"alibaba_interact_sensor_gutil_response,omitempty"` 
    AlibabaInteractSensorGutilResponse
}

/* model for simplify = false
type AlibabaInteractSensorGutilResponse struct {

    // result=0 表示成功
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorGutilResponse struct {

    // result=0 表示成功
    Result   string `json:"result,omitempty"`

}
