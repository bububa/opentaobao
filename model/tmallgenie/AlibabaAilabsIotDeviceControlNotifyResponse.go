package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵IoT异步控制回调接口 APIResponse
alibaba.ailabs.iot.device.control.notify

用于天猫精灵IoT云云接入控制结果的异步回调
*/
type AlibabaAilabsIotDeviceControlNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotDeviceControlNotifyResponse
}

type AlibabaAilabsIotDeviceControlNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_device_control_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否通知成功
    
    RetValue   bool `json:"ret_value,omitempty" xml:"ret_value,omitempty"`

    
}
