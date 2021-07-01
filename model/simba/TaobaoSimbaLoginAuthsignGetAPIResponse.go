package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆权限签名 API返回值 
taobao.simba.login.authsign.get

获取登陆权限签名
*/
type TaobaoSimbaLoginAuthsignGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaLoginAuthsignGetAPIResponseModel
}

// 获取登陆权限签名 成功返回结果
type TaobaoSimbaLoginAuthsignGetAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_login_authsign_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 登陆签名
    SubwayToken   string `json:"subway_token,omitempty" xml:"subway_token,omitempty"`
}
