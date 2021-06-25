package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射接口 APIResponse
taobao.qimen.itemmapping.create

前后端商品映射
*/
type TaobaoQimenItemmappingCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenItemmappingCreateResponse `json:"taobao_qimen_itemmapping_create_response,omitempty"`
}

type TaobaoQimenItemmappingCreateResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
