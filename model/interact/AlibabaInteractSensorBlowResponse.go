package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
吹气 APIResponse
alibaba.interact.sensor.blow

客户端吹气
*/
type AlibabaInteractSensorBlowAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_blow_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0 表示成功
    
    Result   string `json:"result,omitempty" xml:"