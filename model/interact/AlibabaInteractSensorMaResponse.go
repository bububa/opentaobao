package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
码相关API APIResponse
alibaba.interact.sensor.ma

码相关API
*/
type AlibabaInteractSensorMaAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorMaResponse `json:"alibaba_interact_sensor_ma_response,omitempty"`
}

type AlibabaInteractSensorMaResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
