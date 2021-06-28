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
    // Response *TaobaoQimenItemmappingCreateResponse `json:"qimen_itemmapping_create_response,omitempty"` 
    TaobaoQimenItemmappingCreateResponse
}

/* model for simplify = false
type TaobaoQimenItemmappingCreateResponse struct {

    // 
    
    Response  *struct {
        ResponseDO  *ResponseDO `json:"response_do,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenItemmappingCreateResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
