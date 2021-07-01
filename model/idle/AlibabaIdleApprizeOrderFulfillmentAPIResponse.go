package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定担保资金订单履约 API返回值 
alibaba.idle.apprize.order.fulfillment

服务商针对自己的服务订单进行履约
*/
type AlibabaIdleApprizeOrderFulfillmentAPIResponse struct {
    model.CommonResponse
    AlibabaIdleApprizeOrderFulfillmentAPIResponseModel
}

// 鉴定担保资金订单履约 成功返回结果
type AlibabaIdleApprizeOrderFulfillmentAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_idle_apprize_order_fulfillment_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 支付宝流水订单号
    AlipayOrderId   string `json:"alipay_order_id,omitempty" xml:"alipay_order_id,omitempty"`
}
