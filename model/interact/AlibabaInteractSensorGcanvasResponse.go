package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
gcanvas APIResponse
alibaba.interact.sensor.gcanvas

gcanvas 功能
*/
type AlibabaInteractSensorGcanvasAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorGcanvasResponse
}

type AlibabaInteractSensorGcanvasResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_gcanvas_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result=0
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
