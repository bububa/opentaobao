package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川手淘信任登录 APIResponse
taobao.baichuan.user.loginbytoken

百川手淘信任登录
*/
type TaobaoBaichuanUserLoginbytokenAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanUserLoginbytokenResponse
}

type TaobaoBaichuanUserLoginbytokenResponse struct {
    XMLName xml.Name `xml:"baichuan_user_loginbytoken_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
