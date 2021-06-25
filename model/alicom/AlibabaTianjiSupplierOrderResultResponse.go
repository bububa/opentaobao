package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商处理订单接口（订购成功/失败、发货） APIResponse
alibaba.tianji.supplier.order.result

供应商处理订单接口（订购成功/失败、发货）
*/
type AlibabaTianjiSupplierOrderResultAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTianjiSupplierOrderResultResponse `json:"alibaba_tianji_supplier_order_result_response,omitempty"`
}

type AlibabaTianjiSupplierOrderResultResponse struct {

    // 结果
    Model   bool `json:"model,omitempty"`

}
