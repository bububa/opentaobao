package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 APIResponse
taobao.qimen.deliveryorder.batchconfirm

taobao.qimen.deliveryorder.batchconfirm
*/
type TaobaoQimenDeliveryorderBatchconfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenDeliveryorderBatchconfirmResponse `json:"taobao_qimen_deliveryorder_batchconfirm_response,omitempty"`
}

type TaobaoQimenDeliveryorderBatchconfirmResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
