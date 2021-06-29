package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单履约V1 APIResponse
alibaba.idle.recycle.order.fulfillment

外部回收商针对自有回收订单的履行
*/
type AlibabaIdleRecycleOrderFulfillmentAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleOrderFulfillmentResponse
}

type AlibabaIdleRecycleOrderFulfillmentResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_order_fulfillment_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaIdleRecycleOrderFulfillmentResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
