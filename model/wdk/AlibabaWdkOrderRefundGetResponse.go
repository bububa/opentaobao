package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单退款按ID查询 APIResponse
alibaba.wdk.order.refund.get

按照退款ID或者五道口中台订单ID查询退款信息详情
*/
type AlibabaWdkOrderRefundGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOrderRefundGetResponse `json:"alibaba_wdk_order_refund_get_response,omitempty"` 
    AlibabaWdkOrderRefundGetResponse
}

/* model for simplify = false
type AlibabaWdkOrderRefundGetResponse struct {

    // 结果
    
    Result  *struct {
        OrderSyncRefundListResult  *OrderSyncRefundListResult `json:"order_sync_refund_list_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOrderRefundGetResponse struct {

    // 结果
    Result   *OrderSyncRefundListResult `json:"result,omitempty"`

}
