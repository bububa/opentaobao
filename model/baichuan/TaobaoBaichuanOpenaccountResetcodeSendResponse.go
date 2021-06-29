package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川发送找回密码验证码 APIResponse
taobao.baichuan.openaccount.resetcode.send

百川发送找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeSendAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountResetcodeSendResponse
}

type TaobaoBaichuanOpenaccountResetcodeSendResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_resetcode_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
