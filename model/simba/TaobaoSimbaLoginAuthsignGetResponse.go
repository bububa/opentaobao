package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆权限签名 APIResponse
taobao.simba.login.authsign.get

获取登陆权限签名
*/
type TaobaoSimbaLoginAuthsignGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaLoginAuthsignGetResponse `json:"taobao_simba_login_authsign_get_response,omitempty"`
}

type TaobaoSimbaLoginAuthsignGetResponse struct {

    // 登陆签名
    SubwayToken   string `json:"subway_token,omitempty"`

}
