package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
网络状态 APIResponse
alibaba.interact.sensor.networkstatus

客户端网络状态
*/
type AlibabaInteractSensorNetworkstatusAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorNetworkstatusResponse
}

type AlibabaInteractSensorNetworkstatusResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_networkstatus_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // return=0表示成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
