package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP开票请求接口 APIResponse
alibaba.einvoice.createreq

ERP发起开票请求
*/
type AlibabaEinvoiceCreatereqAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceCreatereqResponse
}

type AlibabaEinvoiceCreatereqResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_createreq_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 开票信息是否成功接受
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
