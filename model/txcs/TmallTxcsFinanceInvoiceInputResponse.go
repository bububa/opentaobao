package txcs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商发票录入 APIResponse
tmall.txcs.finance.invoice.input

提供天猫超市外部合作商家财务：供应商发票录入
*/
type TmallTxcsFinanceInvoiceInputAPIResponse struct {
    model.CommonResponse
    TmallTxcsFinanceInvoiceInputResponse
}

type TmallTxcsFinanceInvoiceInputResponse struct {
    XMLName xml.Name `xml:"tmall_txcs_finance_invoice_input_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回内容
    
    Result   *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
