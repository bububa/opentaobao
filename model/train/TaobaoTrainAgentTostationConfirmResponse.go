package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下票确认送票至车站服务 APIResponse
taobao.train.agent.tostation.confirm

送票至车站的订单，代理商确认配送到站
*/
type TaobaoTrainAgentTostationConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentTostationConfirmResponse `json:"taobao_train_agent_tostation_confirm_response,omitempty"`
}

type TaobaoTrainAgentTostationConfirmResponse struct {

    // 错误码
    ErrorMsgCode   string `json:"error_msg_code,omitempty"`

    // 扩展参数
    ExtendParams   string `json:"extend_params,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
