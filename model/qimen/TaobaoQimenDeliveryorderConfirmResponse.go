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
    // Response *TaobaoQimenDeliveryorderConfirmResponse `json:"qimen_deliveryorder_confirm_response,omitempty"` 
    TaobaoQimenDeliveryorderConfirmResponse
}

/* model for simplify = false
type TaobaoQimenDeliveryorderConfirmResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenDeliveryorderConfirmResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
