package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据收件人信息查询交易单号接口 APIResponse
taobao.qimen.order.query

WMS 调用该接口，根据收件人信息查询平台交易订单号。
*/
type TaobaoQimenOrderQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenOrderQueryResponse `json:"qimen_order_query_response,omitempty"` 
    TaobaoQimenOrderQueryResponse
}

/* model for simplify = false
type TaobaoQimenOrderQueryResponse struct {

    // 响应
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenOrderQueryResponse struct {

    // 响应
    Response   *Response `json:"response,omitempty"`

}
