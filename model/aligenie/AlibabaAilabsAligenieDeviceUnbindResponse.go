package aligenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备解绑操作接口 APIResponse
alibaba.ailabs.aligenie.device.unbind

让开发者能根据设备ID进行解绑操作的接口
*/
type AlibabaAilabsAligenieDeviceUnbindAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieDeviceUnbindResponse
}

type AlibabaAilabsAligenieDeviceUnbindResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_device_unbind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 解绑是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
