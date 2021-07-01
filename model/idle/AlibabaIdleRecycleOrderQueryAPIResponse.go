package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单查询V1 API返回值 
alibaba.idle.recycle.order.query

查询回收订单
*/
type AlibabaIdleRecycleOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleOrderQueryAPIResponseModel
}

// 闲鱼回收订单查询V1 成功返回结果
type AlibabaIdleRecycleOrderQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_order_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaIdleRecycleOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
