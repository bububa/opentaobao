package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或者新增庭后信息 APIResponse
alibaba.legal.suit.court.after.push

供外部ISV供应商 推送庭后信息给集团诉讼
*/
type AlibabaLegalSuitCourtAfterPushAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitCourtAfterPushResponse
}

type AlibabaLegalSuitCourtAfterPushResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_court_after_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
