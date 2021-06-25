package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
客户端授权页 APIResponse
alibaba.interact.sensor.authorize

客户端授权页
*/
type AlibabaInteractSensorAuthorizeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorAuthorizeResponse `json:"alibaba_interact_sensor_authorize_response,omitempty"`
}

type AlibabaInteractSensorAuthorizeResponse struct {

    // return=0 表示成功
    Result   string `json:"result,omitempty"`

}
