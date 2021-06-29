package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票扣减的接口 APIResponse
alibaba.einvoice.deduct.get

获取历史发票扣减量、每日发票扣减量的接口
*/
type AlibabaEinvoiceDeductGetAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceDeductGetResponse
}

type AlibabaEinvoiceDeductGetResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_deduct_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaEinvoiceDeductGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
