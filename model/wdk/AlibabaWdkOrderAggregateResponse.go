package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达订单按门店机台号聚合查询 APIResponse
alibaba.wdk.order.aggregate

淘鲜达订单按门店机台号聚合查询
*/
type AlibabaWdkOrderAggregateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_order_aggregate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *OrderAggregateQueryResult `json:"result,omitempty" xml:"