package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射查询接口 APIResponse
taobao.qimen.itemmapping.query

前后端商品映射查询接口
*/
type TaobaoQimenItemmappingQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenItemmappingQueryResponse `json:"qimen_itemmapping_query_response,omitempty"` 
    TaobaoQimenItemmappingQueryResponse
}

/* model for simplify = false
type TaobaoQimenItemmappingQueryResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenItemmappingQueryResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
