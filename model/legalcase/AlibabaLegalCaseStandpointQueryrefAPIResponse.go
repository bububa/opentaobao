package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推送口径信息 API返回值 
alibaba.legal.case.standpoint.queryref

查询推送口径信息
*/
type AlibabaLegalCaseStandpointQueryrefAPIResponse struct {
    model.CommonResponse
    AlibabaLegalCaseStandpointQueryrefAPIResponseModel
}

// 查询推送口径信息 成功返回结果
type AlibabaLegalCaseStandpointQueryrefAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_legal_case_standpoint_queryref_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
