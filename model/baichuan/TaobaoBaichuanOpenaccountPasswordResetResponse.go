package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川找回密码 APIResponse
taobao.baichuan.openaccount.password.reset

百川找回密码
*/
type TaobaoBaichuanOpenaccountPasswordResetAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountPasswordResetResponse
}

type TaobaoBaichuanOpenaccountPasswordResetResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_password_reset_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
