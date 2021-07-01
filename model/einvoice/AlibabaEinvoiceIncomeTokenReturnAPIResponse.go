package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传税号token API返回值 
alibaba.einvoice.income.token.return

服务商回传税号token，用来勾选抵扣认证
*/
type AlibabaEinvoiceIncomeTokenReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceIncomeTokenReturnAPIResponseModel
}

// 服务商回传税号token 成功返回结果
type AlibabaEinvoiceIncomeTokenReturnAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_einvoice_income_token_return_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result接口是否调用成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
