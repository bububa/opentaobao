package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川登录二次验证 APIResponse
taobao.baichuan.openaccount.logindoublecheck

百川登录二次验证
*/
type TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBaichuanOpenaccountLogindoublecheckResponse `json:"taobao_baichuan_openaccount_logindoublecheck_response,omitempty"`
}

type TaobaoBaichuanOpenaccountLogindoublecheckResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
