package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单完结 APIResponse
tmall.servicecenter.workcard.complete

工单完结
*/
type TmallServicecenterWorkcardCompleteAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardCompleteResponse `json:"tmall_servicecenter_workcard_complete_response,omitempty"` 
    TmallServicecenterWorkcardCompleteResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardCompleteResponse struct {

    // 响应结果
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardCompleteResponse struct {

    // 响应结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
