package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 APIResponse
taobao.qimen.deliveryorder.confirm

taobao.qimen.deliveryorder.confirm
*/
type TaobaoQimenDeliveryorderConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenDeliveryorderConfirmResponse `json:"taobao_qimen_deliveryorder_confirm_response,omitempty"`
}

type TaobaoQimenDeliveryorderConfirmResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
