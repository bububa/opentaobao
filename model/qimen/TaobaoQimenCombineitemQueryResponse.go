package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合货品关系查询接口 APIResponse
taobao.qimen.combineitem.query

组合货品关系查询
*/
type TaobaoQimenCombineitemQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenCombineitemQueryResponse `json:"qimen_combineitem_query_response,omitempty"` 
    TaobaoQimenCombineitemQueryResponse
}

/* model for simplify = false
type TaobaoQimenCombineitemQueryResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenCombineitemQueryResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
