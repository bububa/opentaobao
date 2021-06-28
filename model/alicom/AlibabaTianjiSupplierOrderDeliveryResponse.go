package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天机供应商发货 APIResponse
alibaba.tianji.supplier.order.delivery

天机供应商发货
*/
type AlibabaTianjiSupplierOrderDeliveryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTianjiSupplierOrderDeliveryResponse `json:"alibaba_tianji_supplier_order_delivery_response,omitempty"` 
    AlibabaTianjiSupplierOrderDeliveryResponse
}

/* model for simplify = false
type AlibabaTianjiSupplierOrderDeliveryResponse struct {

    // 发货是否成功
    
    Model   bool `json:"model,omitempty"`
    

}
*/

type AlibabaTianjiSupplierOrderDeliveryResponse struct {

    // 发货是否成功
    Model   bool `json:"model,omitempty"`

}
