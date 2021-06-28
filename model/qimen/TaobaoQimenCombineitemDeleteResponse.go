package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合货品删除接口 APIResponse
taobao.qimen.combineitem.delete

组合货品删除
*/
type TaobaoQimenCombineitemDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenCombineitemDeleteResponse `json:"qimen_combineitem_delete_response,omitempty"` 
    TaobaoQimenCombineitemDeleteResponse
}

/* model for simplify = false
type TaobaoQimenCombineitemDeleteResponse struct {

    // 
    
    Response  *struct {
        ResponseDO  *ResponseDO `json:"response_do,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenCombineitemDeleteResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
