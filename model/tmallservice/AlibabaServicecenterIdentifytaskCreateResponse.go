package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建核销单 APIResponse
alibaba.servicecenter.identifytask.create

创建核销单
*/
type AlibabaServicecenterIdentifytaskCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaServicecenterIdentifytaskCreateResponse `json:"alibaba_servicecenter_identifytask_create_response,omitempty"` 
    AlibabaServicecenterIdentifytaskCreateResponse
}

/* model for simplify = false
type AlibabaServicecenterIdentifytaskCreateResponse struct {

    // 请求结果
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaServicecenterIdentifytaskCreateResponse struct {

    // 请求结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
