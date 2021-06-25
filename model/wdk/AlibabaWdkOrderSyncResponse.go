package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口外部订单同步 APIResponse
alibaba.wdk.order.sync

外部商户使用自助POS下单订单同步到五道口
*/
type AlibabaWdkOrderSyncAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkOrderSyncResponse `json:"alibaba_wdk_order_sync_response,omitempty"`
}

type AlibabaWdkOrderSyncResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 返回码
    ReturnCode   int64 `json:"return_code,omitempty"`

    // 描述
    Message   string `json:"message,omitempty"`

    // 订单号
    Target   string `json:"target,omitempty"`

}
