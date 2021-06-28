package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务单查询 APIResponse
alibaba.servicecenter.spserviceorder.query

服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
*/
type AlibabaServicecenterSpserviceorderQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaServicecenterSpserviceorderQueryResponse `json:"alibaba_servicecenter_spserviceorder_query_response,omitempty"` 
    AlibabaServicecenterSpserviceorderQueryResponse
}

/* model for simplify = false
type AlibabaServicecenterSpserviceorderQueryResponse struct {

    // 请求结果
    
    Result  *struct {
        AlibabaServicecenterSpserviceorderQueryResult  *AlibabaServicecenterSpserviceorderQueryResult `json:"alibaba_servicecenter_spserviceorder_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaServicecenterSpserviceorderQueryResponse struct {

    // 请求结果
    Result   *AlibabaServicecenterSpserviceorderQueryResult `json:"result,omitempty"`

}
