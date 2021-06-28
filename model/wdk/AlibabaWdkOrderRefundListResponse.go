package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口交易退款批量查询 APIResponse
alibaba.wdk.order.refund.list

按照条件查询退款数据
*/
type AlibabaWdkOrderRefundListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOrderRefundListResponse `json:"alibaba_wdk_order_refund_list_response,omitempty"` 
    AlibabaWdkOrderRefundListResponse
}

/* model for simplify = false
type AlibabaWdkOrderRefundListResponse struct {

    // 结果内容
    
    Result  *struct {
        OrderSyncRefundListResult  *OrderSyncRefundListResult `json:"order_sync_refund_list_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOrderRefundListResponse struct {

    // 结果内容
    Result   *OrderSyncRefundListResult `json:"result,omitempty"`

}
