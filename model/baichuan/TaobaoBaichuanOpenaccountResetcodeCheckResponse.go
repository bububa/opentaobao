package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川验证找回密码验证码 APIResponse
taobao.baichuan.openaccount.resetcode.check

百川验证找回密码验证码
*/
type TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanOpenaccountResetcodeCheckResponse `json:"baichuan_openaccount_resetcode_check_response,omitempty"` 
    TaobaoBaichuanOpenaccountResetcodeCheckResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountResetcodeCheckResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountResetcodeCheckResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
