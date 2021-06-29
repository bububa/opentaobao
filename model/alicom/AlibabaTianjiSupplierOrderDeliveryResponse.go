package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天机供应商发货 API返回值 
alibaba.tianji.supplier.order.delivery

天机供应商发货
*/
type AlibabaTianjiSupplierOrderDeliveryAPIResponse struct {
    model.CommonResponse
    AlibabaTianjiSupplierOrderDeliveryResponse
}

// 天机供应商发货 成功返回结果
type AlibabaTianjiSupplierOrderDeliveryResponse struct {
    XMLName xml.Name `xml:"alibaba_tianji_supplier_order_delivery_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 发货是否成功
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}
