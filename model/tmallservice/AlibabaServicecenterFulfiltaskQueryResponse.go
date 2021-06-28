package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
核销单查询 APIResponse
alibaba.servicecenter.fulfiltask.query

当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
*/
type AlibabaServicecenterFulfiltaskQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaServicecenterFulfiltaskQueryResponse `json:"alibaba_servicecenter_fulfiltask_query_response,omitempty"` 
    AlibabaServicecenterFulfiltaskQueryResponse
}

/* model for simplify = false
type AlibabaServicecenterFulfiltaskQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaServicecenterFulfiltaskQueryResult  *AlibabaServicecenterFulfiltaskQueryResult `json:"alibaba_servicecenter_fulfiltask_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaServicecenterFulfiltaskQueryResponse struct {

    // 接口返回model
    Result   *AlibabaServicecenterFulfiltaskQueryResult `json:"result,omitempty"`

}
