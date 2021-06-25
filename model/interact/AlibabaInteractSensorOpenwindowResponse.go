package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客户端打开新页面 APIResponse
alibaba.interact.sensor.openwindow

客户端打开新页面
*/
type AlibabaInteractSensorOpenwindowAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorOpenwindowResponse `json:"alibaba_interact_sensor_openwindow_response,omitempty"`
}

type AlibabaInteractSensorOpenwindowResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
