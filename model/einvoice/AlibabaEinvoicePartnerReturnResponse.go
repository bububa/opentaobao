package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开票商回传开票结果 APIResponse
alibaba.einvoice.partner.return

开票商返回开票结果数据
*/
type AlibabaEinvoicePartnerReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoicePartnerReturnResponse
}

type AlibabaEinvoicePartnerReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_partner_return_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务端接收开票回传数据的结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
