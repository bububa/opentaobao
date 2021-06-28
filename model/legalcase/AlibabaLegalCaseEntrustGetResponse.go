package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
委托 APIResponse
alibaba.legal.case.entrust.get

获取委托案件的基本信息
*/
type AlibabaLegalCaseEntrustGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalCaseEntrustGetResponse `json:"alibaba_legal_case_entrust_get_response,omitempty"` 
    AlibabaLegalCaseEntrustGetResponse
}

/* model for simplify = false
type AlibabaLegalCaseEntrustGetResponse struct {

    // alinkappserver系统返回的通用结果类
    
    Result  *struct {
        ServiceResult  *ServiceResult `json:"service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaLegalCaseEntrustGetResponse struct {

    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty"`

}
