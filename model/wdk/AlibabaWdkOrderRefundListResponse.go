package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口交易退款批量查询 APIResponse
alibaba.wdk.order.refund.list

按照条件查询退款数据
*/
type AlibabaWdkOrderRefundListAPIResponse struct {
    model.CommonResponse
    AlibabaWdkOrderRefundListResponse
}

type AlibabaWdkOrderRefundListResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_order_refund_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果内容
    
    Result   *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
