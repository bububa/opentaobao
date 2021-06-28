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
    // Response *TaobaoQimenDeliveryorderBatchconfirmResponse `json:"qimen_deliveryorder_batchconfirm_response,omitempty"` 
    TaobaoQimenDeliveryorderBatchconfirmResponse
}

/* model for simplify = false
type TaobaoQimenDeliveryorderBatchconfirmResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenDeliveryorderBatchconfirmResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
