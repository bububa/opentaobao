package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取裁判登记信息 APIResponse
alibaba.legal.suit.judgement.get

供ISV供应商获取集团法务系统的裁判登记信息
*/
type AlibabaLegalSuitJudgementGetAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitJudgementGetResponse
}

type AlibabaLegalSuitJudgementGetResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_judgement_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
