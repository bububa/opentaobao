package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
开庭时间变更 APIResponse
alibaba.legal.case.court.time.update

修改案件的开庭时间
*/
type AlibabaLegalCaseCourtTimeUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalCaseCourtTimeUpdateResponse `json:"alibaba_legal_case_court_time_update_response,omitempty"` 
    AlibabaLegalCaseCourtTimeUpdateResponse
}

/* model for simplify = false
type AlibabaLegalCaseCourtTimeUpdateResponse struct {

    // alinkappserver系统返回的通用结果类
    
    Result  *struct {
        ServiceResult  *ServiceResult `json:"service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaLegalCaseCourtTimeUpdateResponse struct {

    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty"`

}
