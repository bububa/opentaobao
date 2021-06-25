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
    Response *TaobaoBaichuanOpenaccountNewlogindoublecheckResponse `json:"taobao_baichuan_openaccount_newlogindoublecheck_response,omitempty"`
}

type TaobaoBaichuanOpenaccountNewlogindoublecheckResponse struct {

    // name
    Name   string `json:"name,omitempty"`

}
