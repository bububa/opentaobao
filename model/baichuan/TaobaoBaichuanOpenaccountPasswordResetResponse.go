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
    Response *TaobaoBaichuanOpenaccountPasswordResetResponse `json:"taobao_baichuan_openaccount_password_reset_response,omitempty"`
}

type TaobaoBaichuanOpenaccountPasswordResetResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
