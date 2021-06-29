package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼验货担保订单详情查询V1 APIResponse
alibaba.idle.appraise.order.query

鉴定商调用该接口获取订单状态
*/
type AlibabaIdleAppraiseOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleAppraiseOrderQueryResponse
}

type AlibabaIdleAppraiseOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_appraise_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaIdleAppraiseOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
