package alitripreceipt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪发票冲红 APIResponse
alitrip.receipt.seller.invoice.red

飞猪发票创建
*/
type AlitripReceiptSellerInvoiceRedAPIResponse struct {
    model.CommonResponse
    AlitripReceiptSellerInvoiceRedResponse
}

type AlitripReceiptSellerInvoiceRedResponse struct {
    XMLName xml.Name `xml:"alitrip_receipt_seller_invoice_red_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
