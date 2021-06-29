package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无忧购入仓模式服务商退款处理接口 API返回值 
alibaba.idle.isv.order.dealrefund

闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款
*/
type AlibabaIdleIsvOrderDealrefundAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvOrderDealrefundResponse
}

// 闲鱼无忧购入仓模式服务商退款处理接口 成功返回结果
type AlibabaIdleIsvOrderDealrefundResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_order_dealrefund_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 退款信息
    Module   *RefundDTO `json:"module,omitempty" xml:"module,omitempty"`
}
