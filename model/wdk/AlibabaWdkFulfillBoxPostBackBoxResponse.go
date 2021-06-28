package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
RT收箱回传 APIResponse
alibaba.wdk.fulfill.box.post.back.box

RT收箱后，信息同步履约，履约同通知UMS 容器管理
*/
type AlibabaWdkFulfillBoxPostBackBoxAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkFulfillBoxPostBackBoxResponse `json:"alibaba_wdk_fulfill_box_post_back_box_response,omitempty"` 
    AlibabaWdkFulfillBoxPostBackBoxResponse
}

/* model for simplify = false
type AlibabaWdkFulfillBoxPostBackBoxResponse struct {

    // fulfillLogisticSingleResult
    
    FulfillLogisticSingleResult  *struct {
        FulfillLogisticDefaultResult  *FulfillLogisticDefaultResult `json:"fulfill_logistic_default_result,omitempty"`
    } `json:"fulfill_logistic_single_result,omitempty"`
    

}
*/

type AlibabaWdkFulfillBoxPostBackBoxResponse struct {

    // fulfillLogisticSingleResult
    FulfillLogisticSingleResult   *FulfillLogisticDefaultResult `json:"fulfill_logistic_single_result,omitempty"`

}
