package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定担保资金订单履约 APIResponse
alibaba.idle.apprize.order.fulfillment

服务商针对自己的服务订单进行履约
*/
type AlibabaIdleApprizeOrderFulfillmentAPIResponse struct {
    model.CommonResponse
    AlibabaIdleApprizeOrderFulfillmentResponse
}

type AlibabaIdleApprizeOrderFulfillmentResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_apprize_order_fulfillment_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 支付宝流水订单号
    
    AlipayOrderId   string `json:"alipay_order_id,omitempty" xml:"alipay_order_id,omitempty"`

    
}
