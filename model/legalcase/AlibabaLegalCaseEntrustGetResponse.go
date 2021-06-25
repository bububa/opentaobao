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
    Response *AlibabaLegalCaseEntrustGetResponse `json:"alibaba_legal_case_entrust_get_response,omitempty"`
}

type AlibabaLegalCaseEntrustGetResponse struct {

    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty"`

}
