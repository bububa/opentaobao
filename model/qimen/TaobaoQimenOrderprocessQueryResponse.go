package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单流水查询接口 APIResponse
taobao.qimen.orderprocess.query

ERP调用订单流水查询接口
*/
type TaobaoQimenOrderprocessQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenOrderprocessQueryResponse `json:"qimen_orderprocess_query_response,omitempty"` 
    TaobaoQimenOrderprocessQueryResponse
}

/* model for simplify = false
type TaobaoQimenOrderprocessQueryResponse struct {

    // 
    
    Response  *struct {
        OrderProcessQueryResponse  *OrderProcessQueryResponse `json:"order_process_query_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenOrderprocessQueryResponse struct {

    // 
    Response   *OrderProcessQueryResponse `json:"response,omitempty"`

}
