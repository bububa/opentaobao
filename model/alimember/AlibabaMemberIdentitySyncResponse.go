package alimember

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员身份信息同步 APIResponse
alibaba.member.identity.sync

会员身份信息同步
*/
type AlibabaMemberIdentitySyncAPIResponse struct {
    model.CommonResponse
    AlibabaMemberIdentitySyncResponse
}

type AlibabaMemberIdentitySyncResponse struct {
    XMLName xml.Name `xml:"alibaba_member_identity_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
