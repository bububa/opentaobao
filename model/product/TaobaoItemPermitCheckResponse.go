package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发品资质校验 APIResponse
taobao.item.permit.check

对淘宝商品发品、编辑前的预校验接口
*/
type TaobaoItemPermitCheckAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemPermitCheckResponse `json:"taobao_item_permit_check_response,omitempty"`
}

type TaobaoItemPermitCheckResponse struct {

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否成功
    Error   bool `json:"error,omitempty"`

    // 错误码
    Errorcode   string `json:"errorcode,omitempty"`

}
