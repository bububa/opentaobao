package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步底账数据 API返回值 
alibaba.tax.invoice.sync.ledger

接收第三方服务（如：票易通）同步过来的底账发票数据
*/
type AlibabaTaxInvoiceSyncLedgerAPIResponse struct {
    model.CommonResponse
    AlibabaTaxInvoiceSyncLedgerResponse
}

// 同步底账数据 成功返回结果
type AlibabaTaxInvoiceSyncLedgerResponse struct {
    XMLName xml.Name `xml:"alibaba_tax_invoice_sync_ledger_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应编码
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    // 异常消息
    ResponseMsg   string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
    // true/false
    Succeeded   bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
