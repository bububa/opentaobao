package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼验货担保服务商订单履约V1 APIResponse
alibaba.idle.appraise.order.perform

闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
*/
type AlibabaIdleAppraiseOrderPerformAPIResponse struct {
    model.CommonResponse
    AlibabaIdleAppraiseOrderPerformResponse
}

type AlibabaIdleAppraiseOrderPerformResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_appraise_order_perform_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaIdleAppraiseOrderPerformResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
