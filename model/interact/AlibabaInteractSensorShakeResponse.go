package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
摇一摇 APIResponse
alibaba.interact.sensor.shake

摇一摇
*/
type AlibabaInteractSensorShakeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorShakeResponse `json:"alibaba_interact_sensor_shake_response,omitempty"`
}

type AlibabaInteractSensorShakeResponse struct {

    // 0
    Result   string `json:"result,omitempty"`

}
