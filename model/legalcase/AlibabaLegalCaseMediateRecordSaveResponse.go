package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增调解结果 APIResponse
alibaba.legal.case.mediate.record.save

增加调解沟通记录
*/
type AlibabaLegalCaseMediateRecordSaveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalCaseMediateRecordSaveResponse `json:"alibaba_legal_case_mediate_record_save_response,omitempty"` 
    AlibabaLegalCaseMediateRecordSaveResponse
}

/* model for simplify = false
type AlibabaLegalCaseMediateRecordSaveResponse struct {

    // alinkappserver系统返回的通用结果类
    
    Result  *struct {
        ServiceResult  *ServiceResult `json:"service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaLegalCaseMediateRecordSaveResponse struct {

    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty"`

}
