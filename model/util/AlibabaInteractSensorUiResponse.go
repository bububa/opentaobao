package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
基本ui操作 APIResponse
alibaba.interact.sensor.ui

Weex 基本UI操作
*/
type AlibabaInteractSensorUiAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorUiResponse `json:"alibaba_interact_sensor_ui_response,omitempty"`
}

type AlibabaInteractSensorUiResponse struct {

    // 仅作客户端鉴权使用，不会发送接收请求
    Unnamed   string `json:"unnamed,omitempty"`

}
