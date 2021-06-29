package icbudropshipping

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba查询订单支付结果 APIResponse
alibaba.order.pay.result.query

alibaba查询订单支付结果
*/
type AlibabaOrderPayResultQueryAPIResponse struct {
    model.CommonResponse
    AlibabaOrderPayResultQueryResponse
}

type AlibabaOrderPayResultQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_order_pay_result_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // pay response
    
    Value   *CashierPayResponse `json:"value,omitempty" xml:"value,omitempty"`

    
}
