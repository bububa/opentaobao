package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝OauthCode颁发 API返回值 
taobao.oauth.code.create

手淘无线开放的oauthCode颁发接口
*/
type TaobaoOauthCodeCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOauthCodeCreateAPIResponseModel
}

// 淘宝OauthCode颁发 成功返回结果
type TaobaoOauthCodeCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"oauth_code_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // mock out params
    Test   int64 `json:"test,omitempty" xml:"test,omitempty"`
}
