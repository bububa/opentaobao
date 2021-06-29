package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推送裁判登记信息给集团法务系统 APIResponse
alibaba.legal.suit.judgement.push

isv推送裁判登记信息给集团法务系统
*/
type AlibabaLegalSuitJudgementPushAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitJudgementPushResponse
}

type AlibabaLegalSuitJudgementPushResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_judgement_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
