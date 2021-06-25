package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建订单并发货 APIResponse
taobao.logistics.consign.order.createandsend

创建物流订单，并发货。
*/
type TaobaoLogisticsConsignOrderCreateandsendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLogisticsConsignOrderCreateandsendResponse `json:"taobao_logistics_consign_order_createandsend_response,omitempty"`
}

type TaobaoLogisticsConsignOrderCreateandsendResponse struct {

    // 结果描述
    ResultDesc   string `json:"result_desc,omitempty"`

    // 订单ID
    OrderId   int64 `json:"order_id,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
