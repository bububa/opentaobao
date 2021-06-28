package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
popwindow APIResponse
alibaba.interact.sensor.popwindow

popwindow
*/
type AlibabaInteractSensorPopwindowAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorPopwindowResponse `json:"alibaba_interact_sensor_popwindow_response,omitempty"` 
    AlibabaInteractSensorPopwindowResponse
}

/* model for simplify = false
type AlibabaInteractSensorPopwindowResponse struct {

    // result=0 表示成功
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorPopwindowResponse struct {

    // result=0 表示成功
    Result   string `json:"result,omitempty"`

}
