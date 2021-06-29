package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回收订单履约V2 APIResponse
alibaba.idle.recycle.order.perform

闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
*/
type AlibabaIdleRecycleOrderPerformAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleOrderPerformResponse
}

type AlibabaIdleRecycleOrderPerformResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_order_perform_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
