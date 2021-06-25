package train

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商出票中 APIResponse
taobao.train.agent.handleticket.confirm

代理商出票中
*/
type TaobaoTrainAgentHandleticketConfirmAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTrainAgentHandleticketConfirmResponse `json:"taobao_train_agent_handleticket_confirm_response,omitempty"`
}

type TaobaoTrainAgentHandleticketConfirmResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误码
    TrainErrorCode   string `json:"train_error_code,omitempty"`

    // 错误信息
    TrainErrorMsg   string `json:"train_error_msg,omitempty"`

    // 暂无
    ExtendParams   string `json:"extend_params,omitempty"`

}
