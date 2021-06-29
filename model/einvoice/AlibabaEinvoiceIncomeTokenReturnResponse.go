package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传税号token APIResponse
alibaba.einvoice.income.token.return

服务商回传税号token，用来勾选抵扣认证
*/
type AlibabaEinvoiceIncomeTokenReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceIncomeTokenReturnResponse
}

type AlibabaEinvoiceIncomeTokenReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_income_token_return_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result接口是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
