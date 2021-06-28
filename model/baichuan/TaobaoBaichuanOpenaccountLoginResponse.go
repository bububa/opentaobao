package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川用户名密码登录 APIResponse
taobao.baichuan.openaccount.login

百川用户名密码登录
*/
type TaobaoBaichuanOpenaccountLoginAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanOpenaccountLoginResponse `json:"baichuan_openaccount_login_response,omitempty"` 
    TaobaoBaichuanOpenaccountLoginResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountLoginResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountLoginResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
