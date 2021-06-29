package alimember

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消确认 APIResponse
alibaba.member.identity.rescindfinish

取消确认
*/
type AlibabaMemberIdentityRescindfinishAPIResponse struct {
    model.CommonResponse
    AlibabaMemberIdentityRescindfinishResponse
}

type AlibabaMemberIdentityRescindfinishResponse struct {
    XMLName xml.Name `xml:"alibaba_member_identity_rescindfinish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
