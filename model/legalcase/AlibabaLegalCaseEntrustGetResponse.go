package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
委托 APIResponse
alibaba.legal.case.entrust.get

获取委托案件的基本信息
*/
type AlibabaLegalCaseEntrustGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_legal_case_entrust_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"