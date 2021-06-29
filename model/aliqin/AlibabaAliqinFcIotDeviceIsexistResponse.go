package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
判断设备是否存在 APIResponse
alibaba.aliqin.fc.iot.device.isexist

判断设备是否存在
*/
type AlibabaAliqinFcIotDeviceIsexistAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotDeviceIsexistResponse
}

type AlibabaAliqinFcIotDeviceIsexistResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_device_isexist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAliqinFcIotDeviceIsexistResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
