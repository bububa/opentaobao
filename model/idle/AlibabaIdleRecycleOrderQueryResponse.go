package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V1 APIResponse
alibaba.idle.recycle.order.query

查询回收订单
*/
type AlibabaIdleRecycleOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleOrderQueryResponse
}

type AlibabaIdleRecycleOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaIdleRecycleOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
