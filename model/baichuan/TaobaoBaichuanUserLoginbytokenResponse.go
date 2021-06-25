package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川手淘信任登录 APIResponse
taobao.baichuan.user.loginbytoken

百川手淘信任登录
*/
type TaobaoBaichuanUserLoginbytokenAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanUserLoginbytokenResponse `json:"taobao_baichuan_user_loginbytoken_response,omitempty"`
}

type TaobaoBaichuanUserLoginbytokenResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
