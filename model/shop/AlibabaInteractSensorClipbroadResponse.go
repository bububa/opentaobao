package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Weex页面设置或读取剪切板 APIResponse
alibaba.interact.sensor.clipbroad

Weex页面设置或读取剪切板
*/
type AlibabaInteractSensorClipbroadAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorClipbroadResponse
}

type AlibabaInteractSensorClipbroadResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_clipbroad_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 客户端鉴权使用，实际不会发送或接收数据
    
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`

    
    // 客户端鉴权使用，实际不会发送或接收数据
    
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`

    
}
