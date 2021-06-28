package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推送口径信息 APIResponse
alibaba.legal.case.standpoint.queryref

查询推送口径信息
*/
type AlibabaLegalCaseStandpointQueryrefAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_legal_case_standpoint_queryref_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"