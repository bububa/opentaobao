package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询推送口径信息 APIResponse
alibaba.legal.case.standpoint.queryref

查询推送口径信息
*/
type AlibabaLegalCaseStandpointQueryrefAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalCaseStandpointQueryrefResponse `json:"alibaba_legal_case_standpoint_queryref_response,omitempty"` 
    AlibabaLegalCaseStandpointQueryrefResponse
}

/* model for simplify = false
type AlibabaLegalCaseStandpointQueryrefResponse struct {

    // alinkappserver系统返回的通用结果类
    
    Result  *struct {
        ServiceResult  *ServiceResult `json:"service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaLegalCaseStandpointQueryrefResponse struct {

    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty"`

}
