package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开庭信息推送接口 APIResponse
alibaba.legal.suit.court.open.push

供ISV推送开庭信息给集团诉讼
*/
type AlibabaLegalSuitCourtOpenPushAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitCourtOpenPushResponse
}

type AlibabaLegalSuitCourtOpenPushResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_court_open_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
