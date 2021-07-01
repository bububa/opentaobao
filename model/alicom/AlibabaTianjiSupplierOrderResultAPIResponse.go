package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商处理订单接口（订购成功/失败、发货） API返回值 
alibaba.tianji.supplier.order.result

供应商处理订单接口（订购成功/失败、发货）
*/
type AlibabaTianjiSupplierOrderResultAPIResponse struct {
    model.CommonResponse
    AlibabaTianjiSupplierOrderResultAPIResponseModel
}

// 供应商处理订单接口（订购成功/失败、发货） 成功返回结果
type AlibabaTianjiSupplierOrderResultAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tianji_supplier_order_result_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}
