package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端打开新页面 APIResponse
alibaba.interact.sensor.openwindow

客户端打开新页面
*/
type AlibabaInteractSensorOpenwindowAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorOpenwindowResponse
}

type AlibabaInteractSensorOpenwindowResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_openwindow_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
