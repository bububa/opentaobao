package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天机供应商发货 APIResponse
alibaba.tianji.supplier.order.delivery

天机供应商发货
*/
type AlibabaTianjiSupplierOrderDeliveryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tianji_supplier_order_delivery_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 发货是否成功
    
    Model   bool `json:"model,omitempty" xml:"