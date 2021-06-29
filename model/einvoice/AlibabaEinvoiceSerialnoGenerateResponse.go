package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取统一开票流水号 APIResponse
alibaba.einvoice.serialno.generate

erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
*/
type AlibabaEinvoiceSerialnoGenerateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceSerialnoGenerateResponse
}

type AlibabaEinvoiceSerialnoGenerateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_serialno_generate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    SerialNo   string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`

    
}
