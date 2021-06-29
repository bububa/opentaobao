package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商处理订单接口（订购成功/失败、发货） APIResponse
alibaba.tianji.supplier.order.result

供应商处理订单接口（订购成功/失败、发货）
*/
type AlibabaTianjiSupplierOrderResultAPIResponse struct {
    model.CommonResponse
    AlibabaTianjiSupplierOrderResultResponse
}

type AlibabaTianjiSupplierOrderResultResponse struct {
    XMLName xml.Name `xml:"alibaba_tianji_supplier_order_result_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
}
