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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_networkstatus_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0表示成功
    
    Result   string `json:"result,omitempty" xml:"