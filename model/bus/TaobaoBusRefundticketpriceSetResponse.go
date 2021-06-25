package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退款申请接口 APIResponse
taobao.bus.refundticketprice.set

汽车票代理商利用该接口申请退票
*/
type TaobaoBusRefundticketpriceSetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusRefundticketpriceSetResponse `json:"taobao_bus_refundticketprice_set_response,omitempty"`
}

type TaobaoBusRefundticketpriceSetResponse struct {

    // 退票成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
