package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
无交易类凭证核销 APIResponse
taobao.vmarket.eticket.flow.consume

无交易类凭证核销
*/
type TaobaoVmarketEticketFlowConsumeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketFlowConsumeResponse `json:"taobao_vmarket_eticket_flow_consume_response,omitempty"`
}

type TaobaoVmarketEticketFlowConsumeResponse struct {

    // 执行成功
    RetCode   int64 `json:"ret_code,omitempty"`

    // 错误提示信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
