package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家提交设备信息 APIResponse
alibaba.aliqin.fc.iot.device.post

物联网商家设备信息录入
*/
type AlibabaAliqinFcIotDevicePostAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotDevicePostResponse
}

type AlibabaAliqinFcIotDevicePostResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_device_post_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcIotDevicePostResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
