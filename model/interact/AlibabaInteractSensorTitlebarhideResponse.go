package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
隐藏titleBar APIResponse
alibaba.interact.sensor.titlebarhide

隐藏titleBar
*/
type AlibabaInteractSensorTitlebarhideAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorTitlebarhideResponse
}

type AlibabaInteractSensorTitlebarhideResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_titlebarhide_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0表示成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
