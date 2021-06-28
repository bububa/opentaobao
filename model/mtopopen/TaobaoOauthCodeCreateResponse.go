package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝OauthCode颁发 APIResponse
taobao.oauth.code.create

手淘无线开放的oauthCode颁发接口
*/
type TaobaoOauthCodeCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOauthCodeCreateResponse
}

type TaobaoOauthCodeCreateResponse struct {
    XMLName xml.Name `xml:"oauth_code_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // mock out params
    
    Test   int64 `json:"test,omitempty" xml:"test,omitempty"`

    
}
