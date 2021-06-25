package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
自助机条形码被动支付 APIResponse
taobao.bus.tvmpayorder.set

汽车票线下自助机条形码支付
*/
type TaobaoBusTvmpayorderSetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusTvmpayorderSetResponse `json:"taobao_bus_tvmpayorder_set_response,omitempty"`
}

type TaobaoBusTvmpayorderSetResponse struct {

    // errorCode  线下扫码支付  错误码
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg 错误信息
    ResultMsg   string `json:"result_msg,omitempty"`

    // success true 成功 false 失败
    IsSuccess   bool `json:"is_success,omitempty"`

    // payTime
    PayTime   string `json:"pay_time,omitempty"`

}
