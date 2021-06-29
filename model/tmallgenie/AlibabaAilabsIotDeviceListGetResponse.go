package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取iot设备列表 APIResponse
alibaba.ailabs.iot.device.list.get

通过此接口获取用户名下的iot设备列表
*/
type AlibabaAilabsIotDeviceListGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsIotDeviceListGetResponse
}

type AlibabaAilabsIotDeviceListGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_iot_device_list_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAilabsIotDeviceListGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
