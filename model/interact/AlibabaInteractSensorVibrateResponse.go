package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客户端震动 APIResponse
alibaba.interact.sensor.vibrate

客户端震动
*/
type AlibabaInteractSensorVibrateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorVibrateResponse `json:"alibaba_interact_sensor_vibrate_response,omitempty"`
}

type AlibabaInteractSensorVibrateResponse struct {

    // result=0表示成功
    Result   string `json:"result,omitempty"`

}
