package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录 API返回值 
taobao.baichuan.user.login

百川H5登录
*/
type TaobaoBaichuanUserLoginAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanUserLoginAPIResponseModel
}

// 百川H5登录 成功返回结果
type TaobaoBaichuanUserLoginAPIResponseModel struct {
    XMLName xml.Name `xml:"baichuan_user_login_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
