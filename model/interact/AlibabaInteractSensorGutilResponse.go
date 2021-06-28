package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
canvas工具包 APIResponse
alibaba.interact.sensor.gutil

canvas工具包
*/
type AlibabaInteractSensorGutilAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorGutilResponse
}

type AlibabaInteractSensorGutilResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_gutil_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0 表示成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
