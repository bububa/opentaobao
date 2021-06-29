package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回传订单退款审核结果 APIResponse
alibaba.einvoice.order.refund.update

ISV回传订单退款审核结果
*/
type AlibabaEinvoiceOrderRefundUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceOrderRefundUpdateResponse
}

type AlibabaEinvoiceOrderRefundUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_order_refund_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
