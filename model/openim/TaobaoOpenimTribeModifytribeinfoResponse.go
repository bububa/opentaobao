package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群信息修改 APIResponse
taobao.openim.tribe.modifytribeinfo

OPENIM群信息修改
*/
type TaobaoOpenimTribeModifytribeinfoAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimTribeModifytribeinfoResponse `json:"taobao_openim_tribe_modifytribeinfo_response,omitempty"`
}

type TaobaoOpenimTribeModifytribeinfoResponse struct {

    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty"`

}
