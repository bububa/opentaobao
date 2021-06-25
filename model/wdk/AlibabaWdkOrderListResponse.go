package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单拉取 APIResponse
alibaba.wdk.order.list

五道口交易订单拉取接口
*/
type AlibabaWdkOrderListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkOrderListResponse `json:"alibaba_wdk_order_list_response,omitempty"`
}

type AlibabaWdkOrderListResponse struct {

    // 返回数据
    Result   *AlibabaWdkOrderListResult `json:"result,omitempty"`

}
