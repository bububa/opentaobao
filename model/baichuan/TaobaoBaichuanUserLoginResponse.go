package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录 APIResponse
taobao.baichuan.user.login

百川H5登录
*/
type TaobaoBaichuanUserLoginAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanUserLoginResponse
}

type TaobaoBaichuanUserLoginResponse struct {
    XMLName xml.Name `xml:"baichuan_user_login_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
