package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆权限签名 APIResponse
taobao.simba.login.authsign.get

获取登陆权限签名
*/
type TaobaoSimbaLoginAuthsignGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_login_authsign_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 登陆签名
    
    SubwayToken   string `json:"subway_token,omitempty" xml:"