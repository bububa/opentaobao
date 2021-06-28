package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
gcanvas APIResponse
alibaba.interact.sensor.gcanvas

gcanvas 功能
*/
type AlibabaInteractSensorGcanvasAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorGcanvasResponse `json:"alibaba_interact_sensor_gcanvas_response,omitempty"` 
    AlibabaInteractSensorGcanvasResponse
}

/* model for simplify = false
type AlibabaInteractSensorGcanvasResponse struct {

    // result=0
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorGcanvasResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
