package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
基本ui操作 APIResponse
alibaba.interact.sensor.ui

Weex 基本UI操作
*/
type AlibabaInteractSensorUiAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorUiResponse
}

type AlibabaInteractSensorUiResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_ui_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 仅作客户端鉴权使用，不会发送接收请求
    
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`

    
}
