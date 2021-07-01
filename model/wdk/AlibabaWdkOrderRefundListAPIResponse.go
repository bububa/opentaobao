package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口交易退款批量查询 API返回值 
alibaba.wdk.order.refund.list

按照条件查询退款数据
*/
type AlibabaWdkOrderRefundListAPIResponse struct {
    model.CommonResponse
    AlibabaWdkOrderRefundListAPIResponseModel
}

// 五道口交易退款批量查询 成功返回结果
type AlibabaWdkOrderRefundListAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_order_refund_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果内容
    Result   *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`
}
