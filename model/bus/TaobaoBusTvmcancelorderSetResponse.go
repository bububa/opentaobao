package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机未付款取消订单 APIResponse
taobao.bus.tvmcancelorder.set

自助机汽车票未付款取消订单
*/
type TaobaoBusTvmcancelorderSetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusTvmcancelorderSetResponse `json:"taobao_bus_tvmcancelorder_set_response,omitempty"`
}

type TaobaoBusTvmcancelorderSetResponse struct {

    // 错误码  ORDER_NOT_FOUND
    ResultCode   string `json:"result_code,omitempty"`

    // 错误描述  订单无法查询
    ResultMsg   string `json:"result_msg,omitempty"`

    // true代表成功 false 代表失败
    IsSuccess   bool `json:"is_success,omitempty"`

}
