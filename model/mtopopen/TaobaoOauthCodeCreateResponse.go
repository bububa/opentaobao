package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝OauthCode颁发 APIResponse
taobao.oauth.code.create

手淘无线开放的oauthCode颁发接口
*/
type TaobaoOauthCodeCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOauthCodeCreateResponse `json:"taobao_oauth_code_create_response,omitempty"`
}

type TaobaoOauthCodeCreateResponse struct {

    // mock out params
    Test   int64 `json:"test,omitempty"`

}
