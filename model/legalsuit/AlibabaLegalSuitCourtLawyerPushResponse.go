package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推荐律师接口 APIResponse
alibaba.legal.suit.court.lawyer.push

为诉讼系统推荐律师
*/
type AlibabaLegalSuitCourtLawyerPushAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitCourtLawyerPushResponse
}

type AlibabaLegalSuitCourtLawyerPushResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_court_lawyer_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
