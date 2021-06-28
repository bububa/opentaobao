package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
交易订单详情查询 APIResponse
alibaba.wdk.order.get

五道口三江单据查询接口
*/
type AlibabaWdkOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOrderGetResponse `json:"alibaba_wdk_order_get_response,omitempty"` 
    AlibabaWdkOrderGetResponse
}

/* model for simplify = false
type AlibabaWdkOrderGetResponse struct {

    // 返回数据
    
    Result  *struct {
        AlibabaWdkOrderGetResult  *AlibabaWdkOrderGetResult `json:"alibaba_wdk_order_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOrderGetResponse struct {

    // 返回数据
    Result   *AlibabaWdkOrderGetResult `json:"result,omitempty"`

}
