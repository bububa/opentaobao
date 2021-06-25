package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发票回调接口 APIResponse
taobao.bus.invoice.return

汽车票发票回调接口
*/
type TaobaoBusInvoiceReturnAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusInvoiceReturnResponse `json:"taobao_bus_invoice_return_response,omitempty"`
}

type TaobaoBusInvoiceReturnResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误信息
    ResultMsg   string `json:"result_msg,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
