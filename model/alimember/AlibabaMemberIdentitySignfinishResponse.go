package alimember

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
签约确认 APIResponse
alibaba.member.identity.signfinish

签约确认
*/
type AlibabaMemberIdentitySignfinishAPIResponse struct {
    model.CommonResponse
    AlibabaMemberIdentitySignfinishResponse
}

type AlibabaMemberIdentitySignfinishResponse struct {
    XMLName xml.Name `xml:"alibaba_member_identity_signfinish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
