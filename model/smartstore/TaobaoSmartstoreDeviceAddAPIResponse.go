package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智慧门店设备创建 API返回值 
taobao.smartstore.device.add

智慧门店设备创建
*/
type TaobaoSmartstoreDeviceAddAPIResponse struct {
    model.CommonResponse
    TaobaoSmartstoreDeviceAddAPIResponseModel
}

// 智慧门店设备创建 成功返回结果
type TaobaoSmartstoreDeviceAddAPIResponseModel struct {
    XMLName xml.Name `xml:"smartstore_device_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 设备号
    DeviceCode   string `json:"device_code,omitempty" xml:"device_code,omitempty"`
}
