package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智慧门店设备创建 APIResponse
taobao.smartstore.device.add

智慧门店设备创建
*/
type TaobaoSmartstoreDeviceAddAPIResponse struct {
    model.CommonResponse
    TaobaoSmartstoreDeviceAddResponse
}

type TaobaoSmartstoreDeviceAddResponse struct {
    XMLName xml.Name `xml:"smartstore_device_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 设备号
    
    DeviceCode   string `json:"device_code,omitempty" xml:"device_code,omitempty"`

    
}
