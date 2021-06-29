package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
takePhoto API返回值 
alibaba.interact.sensor.takephoto

客户端takePhoto
*/
type AlibabaInteractSensorTakephotoAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorTakephotoResponse
}

// takePhoto 成功返回结果
type AlibabaInteractSensorTakephotoResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_takephoto_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // return=0表示成功
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
