package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建接口 APIResponse
taobao.qimen.deliveryorder.create

taobao.qimen.deliveryorder.create
*/
type TaobaoQimenDeliveryorderCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenDeliveryorderCreateResponse `json:"taobao_qimen_deliveryorder_create_response,omitempty"`
}

type TaobaoQimenDeliveryorderCreateResponse struct {

    // 
    Response   *DeliveryOrderCreateResponse `json:"response,omitempty"`

}
