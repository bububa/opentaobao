package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川验证找回密码验证码 APIResponse
taobao.baichuan.openaccount.resetcode.check

百川验证找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountResetcodeCheckResponse
}

type TaobaoBaichuanOpenaccountResetcodeCheckResponse struct {
    XMLName xml.Name `xml:"baichuan_openaccount_resetcode_check_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`

    
}
