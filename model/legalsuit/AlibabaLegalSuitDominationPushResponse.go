package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或者保存管辖信息 APIResponse
alibaba.legal.suit.domination.push

ISV推送管辖信息到诉讼平台
*/
type AlibabaLegalSuitDominationPushAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitDominationPushResponse
}

type AlibabaLegalSuitDominationPushResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_domination_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
