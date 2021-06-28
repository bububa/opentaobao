package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品单仓批次库存查询接口 APIResponse
taobao.qimen.inventorybatch.query

ERP 通过该接口查询指定商品的单仓批次库存
*/
type TaobaoQimenInventorybatchQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenInventorybatchQueryResponse `json:"qimen_inventorybatch_query_response,omitempty"` 
    TaobaoQimenInventorybatchQueryResponse
}

/* model for simplify = false
type TaobaoQimenInventorybatchQueryResponse struct {

    // 响应
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenInventorybatchQueryResponse struct {

    // 响应
    Response   *Response `json:"response,omitempty"`

}
