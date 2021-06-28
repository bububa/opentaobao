package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品查询接口 APIResponse
taobao.qimen.singleitem.query

商品查询接口
*/
type TaobaoQimenSingleitemQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenSingleitemQueryResponse `json:"qimen_singleitem_query_response,omitempty"` 
    TaobaoQimenSingleitemQueryResponse
}

/* model for simplify = false
type TaobaoQimenSingleitemQueryResponse struct {

    // 
    
    Response  *struct {
        ResponseDO  *ResponseDO `json:"response_do,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenSingleitemQueryResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
