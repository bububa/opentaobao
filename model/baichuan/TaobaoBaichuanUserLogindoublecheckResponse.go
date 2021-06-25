package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川H5登录二次验证 APIResponse
taobao.baichuan.user.logindoublecheck

百川H5登录二次验证
*/
type TaobaoBaichuanUserLogindoublecheckAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanUserLogindoublecheckResponse `json:"taobao_baichuan_user_logindoublecheck_response,omitempty"`
}

type TaobaoBaichuanUserLogindoublecheckResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
