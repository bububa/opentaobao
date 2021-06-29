package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
gcanvas API返回值 
alibaba.interact.sensor.gcanvas

gcanvas 功能
*/
type AlibabaInteractSensorGcanvasAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorGcanvasResponse
}

// gcanvas 成功返回结果
type AlibabaInteractSensorGcanvasResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_gcanvas_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result=0
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
