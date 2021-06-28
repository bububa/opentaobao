package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务平台工单创建接口 APIResponse
alibaba.servicecenter.workcard.create

创建服务平台工单
*/
type AlibabaServicecenterWorkcardCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaServicecenterWorkcardCreateResponse `json:"alibaba_servicecenter_workcard_create_response,omitempty"` 
    AlibabaServicecenterWorkcardCreateResponse
}

/* model for simplify = false
type AlibabaServicecenterWorkcardCreateResponse struct {

    // 返回参数
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaServicecenterWorkcardCreateResponse struct {

    // 返回参数
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
