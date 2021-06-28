package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川找回密码 APIResponse
taobao.baichuan.openaccount.password.reset

百川找回密码
*/
type TaobaoBaichuanOpenaccountPasswordResetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanOpenaccountPasswordResetResponse `json:"baichuan_openaccount_password_reset_response,omitempty"` 
    TaobaoBaichuanOpenaccountPasswordResetResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountPasswordResetResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountPasswordResetResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
