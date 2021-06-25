package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群成员退出 APIResponse
taobao.openim.tribe.quit

OPENIM群成员退出
*/
type TaobaoOpenimTribeQuitAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimTribeQuitResponse `json:"taobao_openim_tribe_quit_response,omitempty"`
}

type TaobaoOpenimTribeQuitResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
