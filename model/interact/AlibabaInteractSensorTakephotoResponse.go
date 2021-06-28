package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
takePhoto APIResponse
alibaba.interact.sensor.takephoto

客户端takePhoto
*/
type AlibabaInteractSensorTakephotoAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorTakephotoResponse
}

type AlibabaInteractSensorTakephotoResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_takephoto_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0表示成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
