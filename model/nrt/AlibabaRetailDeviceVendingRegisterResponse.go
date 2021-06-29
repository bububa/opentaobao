package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机设备注册 APIResponse
alibaba.retail.device.vending.register

贩卖机注册
*/
type AlibabaRetailDeviceVendingRegisterAPIResponse struct {
    model.CommonResponse
    AlibabaRetailDeviceVendingRegisterResponse
}

type AlibabaRetailDeviceVendingRegisterResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_device_vending_register_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaRetailDeviceVendingRegisterResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
