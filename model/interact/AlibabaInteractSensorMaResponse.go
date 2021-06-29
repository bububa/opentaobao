package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码相关API APIResponse
alibaba.interact.sensor.ma

码相关API
*/
type AlibabaInteractSensorMaAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorMaResponse
}

type AlibabaInteractSensorMaResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_ma_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result=0
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
