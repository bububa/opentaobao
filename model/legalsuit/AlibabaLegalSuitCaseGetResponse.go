package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取案件信息接口v2版本 API返回值 
alibaba.legal.suit.case.get

获取案件信息
*/
type AlibabaLegalSuitCaseGetAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitCaseGetResponse
}

// 获取案件信息接口v2版本 成功返回结果
type AlibabaLegalSuitCaseGetResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_case_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
