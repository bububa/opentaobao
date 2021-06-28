package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川发送找回密码验证码 APIResponse
taobao.baichuan.openaccount.resetcode.send

百川发送找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeSendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanOpenaccountResetcodeSendResponse `json:"baichuan_openaccount_resetcode_send_response,omitempty"` 
    TaobaoBaichuanOpenaccountResetcodeSendResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountResetcodeSendResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountResetcodeSendResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
