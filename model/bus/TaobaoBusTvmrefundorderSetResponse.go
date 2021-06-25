package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机逆向退款接口 APIResponse
taobao.bus.tvmrefundorder.set

汽车票线下自助机 逆向退票接口；用于已出票完成后，再发起退款（注意这是售后退款，如出票异常但是告诉我们出票成功，后续给客户退款，需要调用这个接口，一般开放给财务。出票过程中的失败，请直接调用出票接口并且传递false标志，我们会自动退款。）
*/
type TaobaoBusTvmrefundorderSetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusTvmrefundorderSetResponse `json:"taobao_bus_tvmrefundorder_set_response,omitempty"`
}

type TaobaoBusTvmrefundorderSetResponse struct {

    // errorCode 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg 错误信息
    ResultMsg   string `json:"result_msg,omitempty"`

    // success true 成功 false 失败
    IsSuccess   bool `json:"is_success,omitempty"`

}
