package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传发票ocr的结果 APIResponse
alibaba.einvoice.income.ocr.return

服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传
*/
type AlibabaEinvoiceIncomeOcrReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceIncomeOcrReturnResponse
}

type AlibabaEinvoiceIncomeOcrReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_income_ocr_return_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口调用结果，true=成功，false=失败，subCode以isp开头时需要服务商重试
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
