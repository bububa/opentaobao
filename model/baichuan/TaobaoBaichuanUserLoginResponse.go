package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录 APIResponse
taobao.baichuan.user.login

百川H5登录
*/
type TaobaoBaichuanUserLoginAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanUserLoginResponse `json:"taobao_baichuan_user_login_response,omitempty"`
}

type TaobaoBaichuanUserLoginResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
