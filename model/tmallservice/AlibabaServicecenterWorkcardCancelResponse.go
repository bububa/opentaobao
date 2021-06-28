package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务平台工单取消接口 APIResponse
alibaba.servicecenter.workcard.cancel

取消服务工单
*/
type AlibabaServicecenterWorkcardCancelAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaServicecenterWorkcardCancelResponse `json:"alibaba_servicecenter_workcard_cancel_response,omitempty"` 
    AlibabaServicecenterWorkcardCancelResponse
}

/* model for simplify = false
type AlibabaServicecenterWorkcardCancelResponse struct {

    // 返回参数
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaServicecenterWorkcardCancelResponse struct {

    // 返回参数
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
