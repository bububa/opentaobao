package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家top回调退款明细 APIResponse
taobao.bus.agent.refundticket.confirm

商家通过top回调告知平台退款明细
*/
type TaobaoBusAgentRefundticketConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusAgentRefundticketConfirmResponse `json:"taobao_bus_agent_refundticket_confirm_response,omitempty"`
}

type TaobaoBusAgentRefundticketConfirmResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 退款回调是否收到
    IsSuccess   bool `json:"is_success,omitempty"`

}
