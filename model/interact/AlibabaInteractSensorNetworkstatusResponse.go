package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
网络状态 APIResponse
alibaba.interact.sensor.networkstatus

客户端网络状态
*/
type AlibabaInteractSensorNetworkstatusAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorNetworkstatusResponse `json:"alibaba_interact_sensor_networkstatus_response,omitempty"`
}

type AlibabaInteractSensorNetworkstatusResponse struct {

    // return=0表示成功
    Result   string `json:"result,omitempty"`

}
