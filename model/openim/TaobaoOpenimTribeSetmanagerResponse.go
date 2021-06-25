package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群设置管理员 APIResponse
taobao.openim.tribe.setmanager

OPENIM群设置管理员
*/
type TaobaoOpenimTribeSetmanagerAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimTribeSetmanagerResponse `json:"taobao_openim_tribe_setmanager_response,omitempty"`
}

type TaobaoOpenimTribeSetmanagerResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
