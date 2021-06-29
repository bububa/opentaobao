package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
进项扫描状态回传 APIResponse
alibaba.einvoice.income.scan.return

回传进项扫描每个阶段的状态，比如ocr开始，ocr结束，查验开始，查验结束等
*/
type AlibabaEinvoiceIncomeScanReturnAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceIncomeScanReturnResponse
}

type AlibabaEinvoiceIncomeScanReturnResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_income_scan_return_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否回传成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
