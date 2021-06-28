package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
百川新登录二次验证 APIResponse
taobao.baichuan.openaccount.newlogindoublecheck

百川新登录二次验证
*/
type TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaichuanOpenaccountNewlogindoublecheckResponse `json:"baichuan_openaccount_newlogindoublecheck_response,omitempty"` 
    TaobaoBaichuanOpenaccountNewlogindoublecheckResponse
}

/* model for simplify = false
type TaobaoBaichuanOpenaccountNewlogindoublecheckResponse struct {

    // name
    
    Name   string `json:"name,omitempty"`
    

}
*/

type TaobaoBaichuanOpenaccountNewlogindoublecheckResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
