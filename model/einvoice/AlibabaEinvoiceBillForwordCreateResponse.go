package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码开票结算单同步前开发票 APIResponse
alibaba.einvoice.bill.forword.create

扫码开票结算单同步前开发票，会将数据同步到结算单中
*/
type AlibabaEinvoiceBillForwordCreateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceBillForwordCreateResponse
}

type AlibabaEinvoiceBillForwordCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_bill_forword_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // errorCode
    
    RetCode   string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // errorMessage
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`

    
    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
