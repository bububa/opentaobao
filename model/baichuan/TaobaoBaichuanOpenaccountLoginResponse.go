package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川用户名密码登录 API返回值 
taobao.baichuan.openaccount.login

百川用户名密码登录
*/
type TaobaoBaichuanOpenaccountLoginAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountLoginResponse
}

// 百川用户名密码登录 成功返回结果
type TaobaoBaichuanOpenaccountLoginResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_login_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
