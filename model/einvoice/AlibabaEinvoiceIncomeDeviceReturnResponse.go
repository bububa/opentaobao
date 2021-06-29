package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传客户端设备列表 APIResponse
alibaba.einvoice.income.device.return

服务商回传客户端agent所处环境的设备列表，比如扫描仪
*/
type AlibabaEinvoiceIncomeDeviceReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceIncomeDeviceReturnResponse
}

type AlibabaEinvoiceIncomeDeviceReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_income_device_return_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
