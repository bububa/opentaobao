package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴权 APIResponse
alibaba.guard.access.auth

刷卡鉴权
*/
type AlibabaGuardAccessAuthAPIResponse struct {
    model.CommonResponse
    AlibabaGuardAccessAuthResponse
}

type AlibabaGuardAccessAuthResponse struct {
    XMLName xml.Name `xml:"alibaba_guard_access_auth_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
