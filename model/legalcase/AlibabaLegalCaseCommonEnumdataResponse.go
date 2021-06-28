package legalcase

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取通用枚举值接口 APIResponse
alibaba.legal.case.common.enumdata

获取通用枚举值接口
*/
type AlibabaLegalCaseCommonEnumdataAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLegalCaseCommonEnumdataResponse `json:"alibaba_legal_case_common_enumdata_response,omitempty"` 
    AlibabaLegalCaseCommonEnumdataResponse
}

/* model for simplify = false
type AlibabaLegalCaseCommonEnumdataResponse struct {

    // alinkappserver系统返回的通用结果类
    
    Result  *struct {
        ServiceResult  *ServiceResult `json:"service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaLegalCaseCommonEnumdataResponse struct {

    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty"`

}
