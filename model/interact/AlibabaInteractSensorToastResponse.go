package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
toast APIResponse
alibaba.interact.sensor.toast

toast提示
*/
type AlibabaInteractSensorToastAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorToastResponse `json:"alibaba_interact_sensor_toast_response,omitempty"`
}

type AlibabaInteractSensorToastResponse struct {

    // result=0 表示成功
    Result   string `json:"result,omitempty"`

}
