package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开票商回传开票结果 API返回值 
alibaba.einvoice.partner.return

开票商返回开票结果数据
*/
type AlibabaEinvoicePartnerReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoicePartnerReturnResponse
}

// 开票商回传开票结果 成功返回结果
type AlibabaEinvoicePartnerReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_partner_return_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 服务端接收开票回传数据的结果
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
