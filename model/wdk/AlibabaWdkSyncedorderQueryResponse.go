package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口查询同步订单 APIResponse
alibaba.wdk.syncedorder.query

外部商户查询同步到五道口的订单
*/
type AlibabaWdkSyncedorderQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSyncedorderQueryResponse `json:"alibaba_wdk_syncedorder_query_response,omitempty"` 
    AlibabaWdkSyncedorderQueryResponse
}

/* model for simplify = false
type AlibabaWdkSyncedorderQueryResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 返回码
    
    ReturnCode   int64 `json:"return_code,omitempty"`
    

    // 描述
    
    Message   string `json:"message,omitempty"`
    

    // 订单号
    
    BizOrderId   string `json:"biz_order_id,omitempty"`
    

}
*/

type AlibabaWdkSyncedorderQueryResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 返回码
    ReturnCode   int64 `json:"return_code,omitempty"`

    // 描述
    Message   string `json:"message,omitempty"`

    // 订单号
    BizOrderId   string `json:"biz_order_id,omitempty"`

}
