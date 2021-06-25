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
    Response *TaobaoQimenOrderprocessQueryResponse `json:"taobao_qimen_orderprocess_query_response,omitempty"`
}

type TaobaoQimenOrderprocessQueryResponse struct {

    // 
    Response   *OrderProcessQueryResponse `json:"response,omitempty"`

}
