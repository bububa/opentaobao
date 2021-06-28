package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达退款单按门店聚合查询 APIResponse
alibaba.wdk.refund.aggregate

淘鲜达退款单按门店聚合查询
*/
type AlibabaWdkRefundAggregateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkRefundAggregateResponse
}

type AlibabaWdkRefundAggregateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_refund_aggregate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *RefundAggregateQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
