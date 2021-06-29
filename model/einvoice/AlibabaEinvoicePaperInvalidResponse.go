package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
纸票作废接口 APIResponse
alibaba.einvoice.paper.invalid

作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
*/
type AlibabaEinvoicePaperInvalidAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoicePaperInvalidResponse
}

type AlibabaEinvoicePaperInvalidResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_paper_invalid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口调用是否成功，操作结果tmc异步返回alibaba_invoice_PaperOpsReturn
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
