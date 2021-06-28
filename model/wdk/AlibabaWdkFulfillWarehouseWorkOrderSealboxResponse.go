package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓封箱回告 APIResponse
alibaba.wdk.fulfill.warehouse.work.order.sealbox

仓封箱回告箱与包裹的关系
*/
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkFulfillWarehouseWorkOrderSealboxResponse `json:"alibaba_wdk_fulfill_warehouse_work_order_sealbox_response,omitempty"` 
    AlibabaWdkFulfillWarehouseWorkOrderSealboxResponse
}

/* model for simplify = false
type AlibabaWdkFulfillWarehouseWorkOrderSealboxResponse struct {

    // 失败返回原因
    
    RespMessage   string `json:"resp_message,omitempty"`
    

    // 成功失败码
    
    RespCode   string `json:"resp_code,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AlibabaWdkFulfillWarehouseWorkOrderSealboxResponse struct {

    // 失败返回原因
    RespMessage   string `json:"resp_message,omitempty"`

    // 成功失败码
    RespCode   string `json:"resp_code,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
