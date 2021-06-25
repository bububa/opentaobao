package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下票送票至车站代理商确认用户已取票服务 APIResponse
taobao.train.agent.tostation.receive

送票至车站的订单，代理商确认用户已取票
*/
type TaobaoTrainAgentTostationReceiveAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentTostationReceiveResponse `json:"taobao_train_agent_tostation_receive_response,omitempty"`
}

type TaobaoTrainAgentTostationReceiveResponse struct {

    // 错误码
    ErrorMsgCode   string `json:"error_msg_code,omitempty"`

    // 扩展参数
    ExtendParams   string `json:"extend_params,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
