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
    // Response *TaobaoBaichuanOpenaccountLogindoublecheckResponse `json:"baichuan_openaccount_logindoublecheck_response,omitempty"` 
    TaobaoBaichuanOpenaccountLogindoublecheckResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountLogindoublecheckResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountLogindoublecheckResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
