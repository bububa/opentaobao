package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建群 APIResponse
taobao.openim.tribe.create

创建一个openim的群
*/
type TaobaoOpenimTribeCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimTribeCreateResponse `json:"taobao_openim_tribe_create_response,omitempty"`
}

type TaobaoOpenimTribeCreateResponse struct {

    // 创建群的信息
    TribeInfo   *TribeInfo `json:"tribe_info,omitempty"`

}
