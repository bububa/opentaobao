package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取案件信息接口v2版本 APIResponse
alibaba.legal.suit.case.get

获取案件信息
*/
type AlibabaLegalSuitCaseGetAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitCaseGetResponse
}

type AlibabaLegalSuitCaseGetResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_case_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
