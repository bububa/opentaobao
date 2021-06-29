package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自研ERP开票结果获取 API返回值 
alibaba.einvoice.merchant.result.get

商家自研ERP开票结果获取
*/
type AlibabaEinvoiceMerchantResultGetAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceMerchantResultGetResponse
}

// 商家自研ERP开票结果获取 成功返回结果
type AlibabaEinvoiceMerchantResultGetResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_merchant_result_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 开票返回结果数据列表
    InvoiceResultList   []InvoiceResult `json:"invoice_result_list,omitempty" xml:"invoice_result_list>invoice_result,omitempty"`
}
