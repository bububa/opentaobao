package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发货单创建批量接口 APIResponse
taobao.qimen.deliveryorder.batchcreate

ERP调用接口，将发货信息批量推送给WMS
*/
type TaobaoQimenDeliveryorderBatchcreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenDeliveryorderBatchcreateResponse `json:"qimen_deliveryorder_batchcreate_response,omitempty"` 
    TaobaoQimenDeliveryorderBatchcreateResponse
}

/* model for simplify = false
type TaobaoQimenDeliveryorderBatchcreateResponse struct {

    // 
    
    Response  *struct {
        DeliveryOrderBatchCreateResponse  *DeliveryOrderBatchCreateResponse `json:"delivery_order_batch_create_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenDeliveryorderBatchcreateResponse struct {

    // 
    Response   *DeliveryOrderBatchCreateResponse `json:"response,omitempty"`

}
