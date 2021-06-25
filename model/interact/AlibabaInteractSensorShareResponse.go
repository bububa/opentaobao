package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分享 APIResponse
alibaba.interact.sensor.share

客户端分享
*/
type AlibabaInteractSensorShareAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorShareResponse `json:"alibaba_interact_sensor_share_response,omitempty"`
}

type AlibabaInteractSensorShareResponse struct {

    // return=0表示成功
    Result   string `json:"result,omitempty"`

}
