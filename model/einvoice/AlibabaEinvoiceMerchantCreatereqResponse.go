package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自研ERP开票请求接口 API返回值 
alibaba.einvoice.merchant.createreq

商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
*/
type AlibabaEinvoiceMerchantCreatereqAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceMerchantCreatereqResponse
}

// 商家自研ERP开票请求接口 成功返回结果
type AlibabaEinvoiceMerchantCreatereqResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_merchant_createreq_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 开票信息是否成功接受
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
