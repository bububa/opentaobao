package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 APIResponse
alibaba.interact.sensor.glue

视频播放
*/
type AlibabaInteractSensorGlueAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorGlueResponse
}

type AlibabaInteractSensorGlueResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_glue_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result=0
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
