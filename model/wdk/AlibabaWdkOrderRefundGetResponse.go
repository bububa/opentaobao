package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单退款按ID查询 APIResponse
alibaba.wdk.order.refund.get

按照退款ID或者五道口中台订单ID查询退款信息详情
*/
type AlibabaWdkOrderRefundGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_order_refund_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *OrderSyncRefundListResult `json:"result,omitempty" xml:"