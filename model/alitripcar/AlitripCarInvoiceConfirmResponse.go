package alitripcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票确认接口 APIResponse
alitrip.car.invoice.confirm

飞猪发票回调接口
*/
type AlitripCarInvoiceConfirmAPIResponse struct {
    model.CommonResponse
    AlitripCarInvoiceConfirmResponse
}

type AlitripCarInvoiceConfirmResponse struct {
    XMLName xml.Name `xml:"alitrip_car_invoice_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // 错误消息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 错误码
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`

    
}
