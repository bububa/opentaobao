package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开庭时间变更 APIResponse
alibaba.legal.case.court.time.update

修改案件的开庭时间
*/
type AlibabaLegalCaseCourtTimeUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_legal_case_court_time_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"