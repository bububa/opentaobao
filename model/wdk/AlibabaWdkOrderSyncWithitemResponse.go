package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单和商品同步接口 APIResponse
alibaba.wdk.order.sync.withitem

轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
*/
type AlibabaWdkOrderSyncWithitemAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOrderSyncWithitemResponse `json:"alibaba_wdk_order_sync_withitem_response,omitempty"` 
    AlibabaWdkOrderSyncWithitemResponse
}

/* model for simplify = false
type AlibabaWdkOrderSyncWithitemResponse struct {

    // 调用返回的结果
    
    Result  *struct {
        AlibabaWdkOrderSyncWithitemApiResult  *AlibabaWdkOrderSyncWithitemApiResult `json:"alibaba_wdk_order_sync_withitem_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOrderSyncWithitemResponse struct {

    // 调用返回的结果
    Result   *AlibabaWdkOrderSyncWithitemApiResult `json:"result,omitempty"`

}
