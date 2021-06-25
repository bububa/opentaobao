package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设置ios证书 APIResponse
taobao.openim.ioscert.production.set

设置ios证书
*/
type TaobaoOpenimIoscertProductionSetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimIoscertProductionSetResponse `json:"taobao_openim_ioscert_production_set_response,omitempty"`
}

type TaobaoOpenimIoscertProductionSetResponse struct {

    // 操作成功
    Code   string `json:"code,omitempty"`

}
